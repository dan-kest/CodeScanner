package repositories

import (
	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/internal/models"
	"github.com/dan-kest/cscanner/internal/repositories/tables"
	"gorm.io/gorm"
)

type repoRepository struct {
	conf   *config.Config
	dbConn *gorm.DB
}

func NewRepoRepository(conf *config.Config, dbConn *gorm.DB) *repoRepository {
	return &repoRepository{
		conf:   conf,
		dbConn: dbConn,
	}
}

func (r *repoRepository) ListRepo(paging *models.Paging) ([]*models.Repo, error) {
	repoList := []*models.Repo{
		{
			ID: 1,
		},
		{
			ID: 2,
		},
	}

	return repoList, nil
}

func (r *repoRepository) ViewRepo(id int) (*models.Repo, error) {
	repo := &models.Repo{
		ID: 1,
	}

	return repo, nil
}

func (r *repoRepository) ScanRepo(id int) error {
	return nil
}

func (r *repoRepository) CreateRepo(repo *models.Repo) error {
	_ = tables.Repository{
		Name: repo.Name,
		URL:  repo.URL,
	}

	return nil
}

func (r *repoRepository) UpdateRepo(id int, repo *models.Repo) error {
	return nil
}

func (r *repoRepository) DeleteRepo(id int) error {
	return nil
}
