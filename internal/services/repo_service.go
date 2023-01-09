package services

import (
	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/internal/constants"
	"github.com/dan-kest/cscanner/internal/interfaces"
	"github.com/dan-kest/cscanner/internal/models"
	"github.com/google/uuid"
)

type RepoService struct {
	conf           *config.Config
	repoRepository interfaces.RepoRepository
	scanRepository interfaces.ScanRepository
}

func NewRepoService(conf *config.Config, repoRepository interfaces.RepoRepository, scanRepository interfaces.ScanRepository) *RepoService {
	return &RepoService{
		conf:           conf,
		repoRepository: repoRepository,
		scanRepository: scanRepository,
	}
}

func (s *RepoService) ListRepo(paging *models.Paging) (*models.RepoPagination, error) {
	repoList, totalCount, err := s.repoRepository.ListRepo(paging)
	if err != nil {
		return nil, err
	}

	repoPagination := &models.RepoPagination{
		Page:        paging.Page,
		ItemPerPage: paging.ItemPerPage,
		TotalCount:  totalCount,
		ItemList:    repoList,
	}

	return repoPagination, nil
}

func (s *RepoService) ViewRepo(id uuid.UUID) (*models.Repo, error) {
	repo, err := s.repoRepository.ViewRepo(id)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (s *RepoService) ScanRepo(id uuid.UUID, scanID uuid.UUID) error {
	task := &models.Task{
		RepositoryID: id,
		ScanID:       scanID,
		Status:       constants.ScanStatusQueued,
	}

	return s.scanRepository.CreateScanHistory(task)
}

func (s *RepoService) CreateRepo(repo *models.Repo) (*uuid.UUID, error) {
	return s.repoRepository.CreateRepo(repo)
}

func (s *RepoService) UpdateRepo(id uuid.UUID, repo *models.Repo) error {
	return s.repoRepository.UpdateRepo(id, repo)
}

func (s *RepoService) DeleteRepo(id uuid.UUID) error {
	return s.repoRepository.DeleteRepo(id)
}
