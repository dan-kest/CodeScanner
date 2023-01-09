package services

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/internal/constants"
	"github.com/dan-kest/cscanner/internal/interfaces"
	"github.com/dan-kest/cscanner/internal/models"
	git "github.com/go-git/go-git/v5"
	"github.com/google/uuid"
)

var (
	// Channel act as a job pool for scan workers
	jobs          chan string
	localRepoPath string
	findingRule   []config.FindingRule
	ignore        string
	wg            sync.WaitGroup
)

type ScanService struct {
	conf           *config.Config
	scanRepository interfaces.ScanRepository
}

func NewScanService(conf *config.Config, scanRepository interfaces.ScanRepository) *ScanService {
	jobs = make(chan string)
	localRepoPath = conf.App.Scan.LocalRepoPath
	findingRule = conf.App.Scan.FindingRule
	ignore = conf.App.Scan.Ignore

	return &ScanService{
		conf:           conf,
		scanRepository: scanRepository,
	}
}

// Run "Queued" task from message queue.
func (s *ScanService) RunTask(task *models.Task) error {
	// Save scan status to "In Progress"
	task.Status = constants.ScanStatusInProgress
	if err := s.scanRepository.CreateScanHistory(task); err != nil {
		return err
	}

	// Default scan status to "Failure" in case of error
	task.Status = constants.ScanStatusFailure

	// Clone/Pull git repository to prepare for scan
	path := s.conf.App.Scan.LocalRepoPath + task.RepositoryIDStr
	if err := cloneOrPullRepo(path, task.URL); err != nil {
		return s.scanRepository.CreateScanHistoryAndResult(task, err.Error())
	}

	// Scan repository
	workerCount := s.conf.App.Scan.WorkerCount
	repo, err := scanRepo(task.RepositoryID, path, workerCount)
	if err != nil {
		return s.scanRepository.CreateScanHistoryAndResult(task, err.Error())
	}

	result, err := json.Marshal(repo)
	if err != nil {
		return s.scanRepository.CreateScanHistoryAndResult(task, err.Error())
	}

	// Save scan result with status "Success"
	task.Status = constants.ScanStatusSuccess

	return s.scanRepository.CreateScanHistoryAndResult(task, string(result))
}

// Run git clone to temporary path. If already cloned, pull instead.
// Limited to public repo only for the time being.
func cloneOrPullRepo(path string, url string) error {
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	if errors.Is(err, git.ErrRepositoryAlreadyExists) {
		gitRepo, err := git.PlainOpen(path)
		if err != nil {
			return err
		}

		workingDir, err := gitRepo.Worktree()
		if err != nil {
			return err
		}

		if err := workingDir.Pull(&git.PullOptions{
			Progress: os.Stdout,
		}); err != nil {
			if errors.Is(err, git.NoErrAlreadyUpToDate) {
				return nil
			}

			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

// Run repository code scan.
func scanRepo(id uuid.UUID, path string, workerCount int) (*models.Repo, error) {
	mu := &sync.Mutex{}
	repo := &models.Repo{
		ID:       id,
		Findings: []*models.Finding{},
	}

	// Start workers
	for w := 1; w <= workerCount; w++ {
		go scanWorker(mu, repo)
	}

	// Start looping files and directories in repository
	if err := loopRepoFiles(path); err != nil {
		return nil, err
	}

	wg.Wait()

	return repo, nil
}

// Recursively loop all files in directories/subdirectories, assign each file path to workers.
func loopRepoFiles(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	directoryList := []string{}
	for _, file := range files {
		// Ignore configured file name
		if strings.Contains(ignore, file.Name()) {
			continue
		}

		fullPath := filepath.Join(path, file.Name())

		if file.IsDir() {
			directoryList = append(directoryList, fullPath)

			continue
		}

		wg.Add(1)
		go func() {
			jobs <- fullPath
		}()
	}

	for _, fullPath := range directoryList {
		if err := loopRepoFiles(fullPath); err != nil {
			return err
		}
	}

	return nil
}

// Worker. Wait for a job from job pool then execute file scan.
// First argument is an instance of Mutex shared among scan workers.
func scanWorker(mu *sync.Mutex, repo *models.Repo) {
	for fullPath := range jobs {
		repoPath := strings.Replace(fullPath, localRepoPath+repo.ID.String(), "", 1)
		findingList, err := scanFile(fullPath, repoPath)
		if err != nil {
			wg.Done()
			log.Panicf("Scam worker error: %s", err)
		}

		mu.Lock()
		repo.Findings = append(repo.Findings, findingList...)
		mu.Unlock()

		wg.Done()
	}
}

// File read & words scan.
func scanFile(fullPath string, repoPath string) ([]*models.Finding, error) {
	findingList := []*models.Finding{}

	file, err := os.Open(filepath.Clean(fullPath))
	defer file.Close() // nolint
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)
	lineCount := 1
	for {
		line, isNotEndOfLine, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		if !isNotEndOfLine {
			lineCount++
		}

		wordList := bytes.Split(line, []byte(constants.ScanWordDelimiter))
		for _, word := range wordList {
			word := strings.TrimSpace(string(word))
			if word != "" {
				finding := matchFindingRule(word)
				if finding != nil {
					finding.Location = models.FindingLocation{
						Path: repoPath,
						Positions: models.FindingLocationPosition{
							Begin: models.FindingLocationPositionBegin{
								Line: lineCount,
							},
						},
					}

					findingList = append(findingList, finding)
				}
			}
		}
	}

	return findingList, nil
}

// Check if input word is matches with any rules, returns models.Finding stuct.
func matchFindingRule(word string) *models.Finding {
	for _, rule := range findingRule {
		isMatch := false
		switch strings.ToLower(rule.Match) {
		case constants.FindingRuleMatchPrefix:
			isMatch = strings.HasPrefix(word, rule.Type)
		case constants.FindingRuleMatchSuffix:
			isMatch = strings.HasSuffix(word, rule.Type)
		case constants.FindingRuleMatchWhole:
			isMatch = word == rule.Type
		case constants.FindingRuleMatchPartial:
			isMatch = strings.Contains(word, rule.Type)
		}

		if isMatch {
			return &models.Finding{
				Type:   rule.Type,
				RuleID: rule.RuleID,
				Metadata: models.FindingMetadata{
					Description: rule.Description,
					Severity:    rule.Severity,
				},
			}
		}
	}

	return nil
}

// Get called only when there's an error in consumer, act as a critical fallback.
func (s *ScanService) RunErrorTask(body []byte, err error) error {
	return s.scanRepository.CreateTaskErrorLog(body, err)
}
