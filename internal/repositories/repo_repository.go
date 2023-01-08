package repositories

import (
	"fmt"

	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/database"
	"github.com/dan-kest/cscanner/internal/constants"
	"github.com/dan-kest/cscanner/internal/models"
	"github.com/dan-kest/cscanner/internal/repositories/tables"
	"github.com/google/uuid"
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

func (r *repoRepository) ListRepo(paging *models.Paging) ([]*models.Repo, int, error) {
	rows := []tables.Repository{}
	var totalCount int64

	tx := database.WithTimeout(r.dbConn)
	tx.Model(&rows).Count(&totalCount)
	tx = buildPaging(tx, r.conf, paging)
	if result := tx.Find(&rows); result.Error != nil {
		return nil, 0, result.Error
	}

	repoList := []*models.Repo{}
	for _, row := range rows {
		repoList = append(repoList, &models.Repo{
			ID:   row.ID,
			Name: row.Name,
			URL:  row.URL,
			// TODO: Map the rest
		})
	}

	return repoList, int(totalCount), nil
}

func (r *repoRepository) ViewRepo(id uuid.UUID) (*models.Repo, error) {
	rows := []tables.Repository{}

	filter := tables.Repository{
		ID: id,
	}

	tx := database.WithTimeout(r.dbConn)

	if result := tx.Find(&rows, &filter); result.Error != nil {
		return nil, result.Error
	}

	if len(rows) == 0 {
		return nil, fmt.Errorf("id %s", constants.ErrorNotFoundSuffix)
	}

	repo := &models.Repo{
		ID:   rows[0].ID,
		Name: rows[0].Name,
		URL:  rows[0].URL,
		// TODO: Map the rest
	}

	return repo, nil
}

func (r *repoRepository) ScanRepo(id uuid.UUID) error {
	// TODO: ???

	return nil
}

func (r *repoRepository) CreateRepo(repo *models.Repo) (*uuid.UUID, error) {
	row := tables.Repository{
		Name: repo.Name,
		URL:  repo.URL,
	}

	tx := database.WithTimeout(r.dbConn)

	if result := tx.Create(&row); result.Error != nil {
		return nil, result.Error
	}

	return &row.ID, nil
}

func (r *repoRepository) UpdateRepo(id uuid.UUID, repo *models.Repo) error {
	filter := tables.Repository{
		ID: id,
	}

	row := tables.Repository{
		Name: repo.Name,
		URL:  repo.URL,
	}

	tx := database.WithTimeout(r.dbConn)

	if result := tx.Model(&filter).Updates(row); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repoRepository) DeleteRepo(id uuid.UUID) error {
	filter := tables.Repository{
		ID: id,
	}

	tx := database.WithTimeout(r.dbConn)

	if result := tx.Delete(&filter); result.Error != nil {
		return result.Error
	}

	return nil
}
