package interfaces

import "github.com/dan-kest/cscanner/internal/models"

type RepoRepository interface {
	ListRepo(paging *models.Paging) ([]*models.Repo, error)
	ViewRepo(id int) (*models.Repo, error)
	CreateRepo(repo *models.Repo) error
	UpdateRepo(id int, repo *models.Repo) error
	DeleteRepo(id int) error
	ScanRepo(id int) error
}
