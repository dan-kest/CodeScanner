package repositories

import (
	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/database"
	"github.com/dan-kest/cscanner/internal/models"
	"github.com/dan-kest/cscanner/internal/repositories/tables"
	"gorm.io/gorm"
)

type scanRepository struct {
	conf   *config.Config
	dbConn *gorm.DB
}

func NewScanRepository(conf *config.Config, dbConn *gorm.DB) *scanRepository {
	return &scanRepository{
		conf:   conf,
		dbConn: dbConn,
	}
}

func (r *scanRepository) CreateScanHistory(scanHistory *models.Task) error {
	row := tables.ScanHistory{
		RepositoryID: scanHistory.RepositoryID,
		ScanID:       scanHistory.ScanID,
		Status:       string(scanHistory.Status),
	}

	tx := database.WithTimeout(r.dbConn)

	if result := tx.Create(&row); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *scanRepository) CreateScanHistoryAndResult(scanHistory *models.Task, result string) error {
	scanHistoryRow := tables.ScanHistory{
		RepositoryID: scanHistory.RepositoryID,
		ScanID:       scanHistory.ScanID,
		Status:       string(scanHistory.Status),
	}

	scanResultRow := tables.ScanResult{
		ID:     scanHistory.ScanID,
		Result: result,
	}

	tx := database.WithTimeout(r.dbConn)

	return tx.Transaction(func(tx *gorm.DB) error {
		if result := tx.Create(&scanHistoryRow); result.Error != nil {
			return result.Error
		}

		if result := tx.Create(&scanResultRow); result.Error != nil {
			return result.Error
		}

		return nil
	})
}

func (r *scanRepository) CreateTaskErrorLog(body []byte, err error) error {
	row := tables.TaskErrorLog{
		Body:    string(body),
		Message: err.Error(),
	}

	tx := database.WithTimeout(r.dbConn)

	if result := tx.Create(&row); result.Error != nil {
		return result.Error
	}

	return nil
}
