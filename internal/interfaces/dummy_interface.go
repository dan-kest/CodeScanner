// nolint:revive
package interfaces

import (
	"errors"

	"github.com/dan-kest/cscanner/internal/models"
	"github.com/google/uuid"
)

/*
	Dummy interface with dummy functions for testing purpose.
*/

type DummyRepoRepositorySuccess struct {
	RepoRepository
}

type DummyRepoRepositoryFail struct {
	RepoRepository
}

func (r *DummyRepoRepositorySuccess) ListRepo(paging *models.Paging) ([]*models.Repo, int, error) {
	return []*models.Repo{}, 1, nil
}

func (r *DummyRepoRepositoryFail) ListRepo(paging *models.Paging) ([]*models.Repo, int, error) {
	return nil, 0, errors.New("")
}

func (r *DummyRepoRepositorySuccess) ViewRepo(id uuid.UUID) (*models.Repo, error) {
	return &models.Repo{}, nil
}

func (r *DummyRepoRepositoryFail) ViewRepo(id uuid.UUID) (*models.Repo, error) {
	return nil, errors.New("")
}

func (r *DummyRepoRepositorySuccess) FetchRepo(id uuid.UUID) (*models.Repo, error) {
	return &models.Repo{}, nil
}

func (r *DummyRepoRepositoryFail) FetchRepo(id uuid.UUID) (*models.Repo, error) {
	return nil, errors.New("")
}

func (r *DummyRepoRepositorySuccess) CreateRepo(repo *models.Repo) (*uuid.UUID, error) {
	id := uuid.New()

	return &id, nil
}

func (r *DummyRepoRepositoryFail) CreateRepo(repo *models.Repo) (*uuid.UUID, error) {
	return nil, errors.New("")
}

func (r *DummyRepoRepositorySuccess) UpdateRepo(id uuid.UUID, repo *models.Repo) error {
	return nil
}

func (r *DummyRepoRepositoryFail) UpdateRepo(id uuid.UUID, repo *models.Repo) error {
	return errors.New("")
}

func (r *DummyRepoRepositorySuccess) DeleteRepo(id uuid.UUID) error {
	return nil
}

func (r *DummyRepoRepositoryFail) DeleteRepo(id uuid.UUID) error {
	return errors.New("")
}

type DummyScanRepositorySuccess struct {
	ScanRepository
}

type DummyScanRepositoryFail struct {
	ScanRepository
}

func (r *DummyScanRepositorySuccess) CreateScanHistory(scanHistory *models.Task) error {
	return nil
}

func (r *DummyScanRepositoryFail) CreateScanHistory(scanHistory *models.Task) error {
	return errors.New("")
}

func (r *DummyScanRepositorySuccess) CreateScanHistoryAndResult(scanHistory *models.Task, result string) error {
	return nil
}

func (r *DummyScanRepositoryFail) CreateScanHistoryAndResult(scanHistory *models.Task, result string) error {
	return errors.New("")
}

func (r *DummyScanRepositorySuccess) CreateTaskErrorLog(body []byte, err error) error {
	return nil
}

func (r *DummyScanRepositoryFail) CreateTaskErrorLog(body []byte, err error) error {
	return errors.New("")
}
