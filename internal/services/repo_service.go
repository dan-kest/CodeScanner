package services

import (
	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/internal/interfaces"
	"github.com/dan-kest/cscanner/internal/models"
)

type RepoService struct {
	conf           *config.Config
	repoRepository interfaces.RepoRepository
}

func NewRepoService(conf *config.Config, repoRepository interfaces.RepoRepository) *RepoService {
	return &RepoService{
		conf:           conf,
		repoRepository: repoRepository,
	}
}

func (s *RepoService) ListRepo(paging *models.Paging) ([]*models.Repo, error) {
	repoList, err := s.repoRepository.ListRepo(paging)
	if err != nil {
		return nil, err
	}

	return repoList, nil
}

func (s *RepoService) ViewRepo(id int) (*models.Repo, error) {
	repo, err := s.repoRepository.ViewRepo(id)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (s *RepoService) ScanRepo(id int) error {
	return s.repoRepository.ScanRepo(id)
}

func (s *RepoService) CreateRepo(repo *models.Repo) error {
	return s.repoRepository.CreateRepo(repo)
}

func (s *RepoService) UpdateRepo(id int, repo *models.Repo) error {
	return s.repoRepository.UpdateRepo(id, repo)
}

func (s *RepoService) DeleteRepo(id int) error {
	return s.repoRepository.DeleteRepo(id)
}
