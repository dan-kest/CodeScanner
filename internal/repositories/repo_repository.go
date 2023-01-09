package repositories

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/database"
	"github.com/dan-kest/cscanner/internal/constants"
	"github.com/dan-kest/cscanner/internal/models"
	"github.com/dan-kest/cscanner/internal/repositories/tables"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	tx = tx.Preload("ScanHistoryList", func(tx *gorm.DB) *gorm.DB {
		return tx.Limit(1).Order(clause.OrderByColumn{
			Column: clause.Column{
				Name: "id",
			},
			Desc: true,
		})
	})
	paging.SetFallback(r.conf.App.Paging)
	tx = buildPaging(tx, paging)
	if result := tx.Find(&rows); result.Error != nil {
		return nil, 0, result.Error
	}

	repoList := []*models.Repo{}
	for _, row := range rows {
		repo := &models.Repo{
			ID: row.ID,
		}
		if row.Name.Valid {
			repo.Name = &row.Name.String
		}
		if row.URL.Valid {
			repo.URL = &row.URL.String
		}

		if len(row.ScanHistoryList) > 0 {
			repo.ScanStatus = constants.ScanStatus(row.ScanHistoryList[0].Status)
			repo.Timestamp = &row.ScanHistoryList[0].CreatedAt
		}

		repoList = append(repoList, repo)
	}

	return repoList, int(totalCount), nil
}

func (r *repoRepository) ViewRepo(id uuid.UUID) (*models.Repo, error) {
	row := tables.Repository{}

	tx := database.WithTimeout(r.dbConn)
	tx = tx.Preload("ScanHistoryList", func(tx *gorm.DB) *gorm.DB {
		return tx.Limit(1).Order(clause.OrderByColumn{
			Column: clause.Column{
				Name: "id",
			},
			Desc: true,
		})
	}).
		Preload("ScanHistoryList.ScanResult")

	if result := tx.First(&row, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("id %s", constants.ErrorNotFoundSuffix)
		}

		return nil, result.Error
	}

	repo := &models.Repo{
		ID: row.ID,
	}
	if row.Name.Valid {
		repo.Name = &row.Name.String
	}
	if row.URL.Valid {
		repo.URL = &row.URL.String
	}

	if len(row.ScanHistoryList) > 0 {
		repo.ScanStatus = constants.ScanStatus(row.ScanHistoryList[0].Status)
		repo.Timestamp = &row.ScanHistoryList[0].CreatedAt

		scanResult := row.ScanHistoryList[0].ScanResult
		if scanResult != nil && scanResult.Result != "" {
			_ = json.Unmarshal([]byte(scanResult.Result), repo)
		}
	}

	return repo, nil
}

func (r *repoRepository) CreateRepo(repo *models.Repo) (*uuid.UUID, error) {
	row := tables.Repository{}
	if repo.Name != nil {
		row.Name.String = *repo.Name
		row.Name.Valid = true
	}
	if repo.URL != nil {
		row.URL.String = *repo.URL
		row.URL.Valid = true
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

	row := tables.Repository{}
	if repo.Name != nil {
		row.Name.String = *repo.Name
		row.Name.Valid = true
	}
	if repo.URL != nil {
		row.URL.String = *repo.URL
		row.URL.Valid = true
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
