package services

import (
	"errors"
	"os"

	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/internal/interfaces"
	"github.com/dan-kest/cscanner/internal/models"
	git "github.com/go-git/go-git/v5"
)

type ScanService struct {
	conf           *config.Config
	scanRepository interfaces.ScanRepository
}

func NewScanService(conf *config.Config, scanRepository interfaces.ScanRepository) *ScanService {
	return &ScanService{
		conf:           conf,
		scanRepository: scanRepository,
	}
}

func (s *ScanService) RunTask(task *models.Task) error {
	path := s.conf.App.Scan.RepoPath + task.RepositoryIDStr
	if err := cloneOrPullRepo(path, task.URL); err != nil {
		return err
	}

	repo := &models.Repo{}
	if err := scanRepo(path, repo); err != nil {
		return err
	}

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
func scanRepo(path string, repo *models.Repo) error {
	return nil
}

func (s *ScanService) RunErrorTask(body []byte, err error) error {
	return s.scanRepository.CreateTaskErrorLog(body, err)
}
