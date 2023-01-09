package services

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/internal/constants"
	"github.com/dan-kest/cscanner/internal/interfaces"
	"github.com/dan-kest/cscanner/internal/models"
	git "github.com/go-git/go-git/v5"
)

var (
	// Channel act as a job pool for scan workers
	jobs          chan string
	wordDelimiter string
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

	// If configured delimiter is invalid (More than 1 byte), fallback to a default delimiter
	wordDelimiter = conf.App.Scan.WordDelimiter
	byteArray := []byte(wordDelimiter)
	if len(byteArray) > 1 {
		wordDelimiter = string(byteArray[:1])
	}

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
	s.scanRepository.CreateScanHistory(task)

	// Default scan status to "Failure" in case of error
	task.Status = constants.ScanStatusFailure

	// Clone/Pull git repository to prepare for scan
	path := s.conf.App.Scan.RepoPath + task.RepositoryIDStr
	if err := cloneOrPullRepo(path, task.URL); err != nil {
		s.scanRepository.CreateScanHistoryAndResult(task, err.Error())

		return err
	}

	// Scan repository
	workerCount := s.conf.App.Scan.WorkerCount
	repo, err := scanRepo(path, workerCount)
	if err != nil {
		s.scanRepository.CreateScanHistoryAndResult(task, err.Error())

		return err
	}

	result, err := json.Marshal(repo)
	if err != nil {
		s.scanRepository.CreateScanHistoryAndResult(task, err.Error())

		return err
	}

	// Save scan result with status "Success"
	task.Status = constants.ScanStatusSuccess
	s.scanRepository.CreateScanHistoryAndResult(task, string(result))

	return nil
}

// Run git clone to temporary path. If already cloned, pull instead.
// Limited to public repo only for the time being.
func cloneOrPullRepo(path string, url string) error {
	gitRepo, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	if errors.Is(err, git.ErrRepositoryAlreadyExists) {
		gitRepo, err = git.PlainOpen(path)
		if err != nil {
			return err
		}

		workingDir, err := gitRepo.Worktree()
		if err != nil {
			return err
		}

		workingDir.Pull(&git.PullOptions{
			Progress: os.Stdout,
		})
	} else if err != nil {
		return err
	}

	return nil
}

// Run repository code scan.
func scanRepo(path string, workerCount int) (*models.Repo, error) {
	var mu *sync.Mutex = &sync.Mutex{}
	repo := &models.Repo{
		Findings: []*models.Finding{},
	}

	// Start workers
	for w := 1; w <= workerCount; w++ {
		go scanWorker(mu, repo)
	}

	// Start looping files and directories in repository
	loopRepoFiles(path)
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

		go func() {
			wg.Add(1)
			jobs <- fullPath
		}()
	}

	for _, fullPath := range directoryList {
		loopRepoFiles(fullPath)
	}

	return nil
}

// Worker. Wait for a job from job pool then execute file scan.
// First argument is an instance of Mutex shared among scan workers.
func scanWorker(mu *sync.Mutex, repo *models.Repo) error {
	for fullPath := range jobs {
		findingList, err := scanFile(fullPath)
		if err != nil {
			wg.Done()
			return err
		}

		mu.Lock()
		repo.Findings = append(repo.Findings, findingList...)
		mu.Unlock()

		wg.Done()
	}

	return nil
}

// File read & words scan.
func scanFile(fullPath string) ([]*models.Finding, error) {
	findingList := []*models.Finding{}

	file, err := os.Open(fullPath)
	defer file.Close()
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
			lineCount += 1
		}

		wordList := bytes.Split(line, []byte(wordDelimiter))
		for _, word := range wordList {
			word := strings.TrimSpace(string(word))
			if word != "" {
				finding := matchFindingRule(word)
				if finding != nil {
					finding.Location = models.FindingLocation{
						Path: fullPath,
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
