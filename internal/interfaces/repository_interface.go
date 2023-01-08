package interfaces

import (
	"github.com/dan-kest/cscanner/internal/models"
	"github.com/google/uuid"
)

type RepoRepository interface {
	ListRepo(paging *models.Paging) ([]*models.Repo, int, error)
	ViewRepo(id uuid.UUID) (*models.Repo, error)
	ScanRepo(id uuid.UUID) error
	CreateRepo(repo *models.Repo) (*uuid.UUID, error)
	UpdateRepo(id uuid.UUID, repo *models.Repo) error
	DeleteRepo(id uuid.UUID) error
}

type ScanRepository interface {
	CreateScanHistory(scanHistory *models.Task) error
	CreateScanHistoryAndResult(scanHistory *models.Task, result string) error
	CreateTaskErrorLog(body []byte, err error) error
}
