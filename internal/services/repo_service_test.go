package services

import (
	"errors"
	"testing"

	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/internal/interfaces"
	"github.com/dan-kest/cscanner/internal/models"
	"github.com/dan-kest/cscanner/pkg/tests"
	"github.com/google/uuid"
)

func TestListRepo(t *testing.T) {
	conf := config.ReadDummy()

	testList := []struct {
		name           string
		paging         *models.Paging
		wantErr        bool
		resultErr      error
		repoRepository interfaces.RepoRepository
	}{
		{
			name:           "TestListRepo::BestCase",
			paging:         &models.Paging{Page: 1, ItemPerPage: 20},
			wantErr:        false,
			resultErr:      nil,
			repoRepository: &interfaces.DummyRepoRepositorySuccess{},
		},
		{
			name:           "TestListRepo::RepoRepositoryError",
			paging:         &models.Paging{Page: 1, ItemPerPage: 20},
			wantErr:        true,
			resultErr:      errors.New(""),
			repoRepository: &interfaces.DummyRepoRepositoryFail{},
		},
	}

	for _, tt := range testList {
		t.Run(tt.name, func(t *testing.T) {
			service := NewRepoService(conf, tt.repoRepository, nil)
			_, err := service.ListRepo(tt.paging)
			if err := tests.CompareError(tt.wantErr, err, tt.resultErr); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestViewRepo(t *testing.T) {
	conf := config.ReadDummy()
	id := uuid.New()

	testList := []struct {
		name           string
		id             uuid.UUID
		wantErr        bool
		resultErr      error
		repoRepository interfaces.RepoRepository
	}{
		{
			name:           "TestViewRepo::BestCase",
			id:             id,
			wantErr:        false,
			resultErr:      nil,
			repoRepository: &interfaces.DummyRepoRepositorySuccess{},
		},
		{
			name:           "TestViewRepo::RepoRepositoryError",
			id:             id,
			wantErr:        true,
			resultErr:      errors.New(""),
			repoRepository: &interfaces.DummyRepoRepositoryFail{},
		},
	}

	for _, tt := range testList {
		t.Run(tt.name, func(t *testing.T) {
			service := NewRepoService(conf, tt.repoRepository, nil)
			_, err := service.ViewRepo(tt.id)
			if err := tests.CompareError(tt.wantErr, err, tt.resultErr); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestScanRepo(t *testing.T) {
	conf := config.ReadDummy()
	id := uuid.New()
	scanID := uuid.New()

	testList := []struct {
		name           string
		id             uuid.UUID
		scanID         uuid.UUID
		wantErr        bool
		resultErr      error
		scanRepository interfaces.ScanRepository
	}{
		{
			name:           "TestScanRepo::BestCase",
			id:             id,
			scanID:         scanID,
			wantErr:        false,
			resultErr:      nil,
			scanRepository: &interfaces.DummyScanRepositorySuccess{},
		},
		{
			name:           "TestListRepo::ScanRepositoryError",
			id:             id,
			scanID:         scanID,
			wantErr:        true,
			resultErr:      errors.New(""),
			scanRepository: &interfaces.DummyScanRepositoryFail{},
		},
	}

	for _, tt := range testList {
		t.Run(tt.name, func(t *testing.T) {
			service := NewRepoService(conf, nil, tt.scanRepository)
			err := service.ScanRepo(tt.id, tt.scanID)
			if err := tests.CompareError(tt.wantErr, err, tt.resultErr); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestCreateRepo(t *testing.T) {
	conf := config.ReadDummy()
	repo := &models.Repo{}

	testList := []struct {
		name           string
		repo           *models.Repo
		wantErr        bool
		resultErr      error
		repoRepository interfaces.RepoRepository
	}{
		{
			name:           "TestCreateRepo::BestCase",
			repo:           repo,
			wantErr:        false,
			resultErr:      nil,
			repoRepository: &interfaces.DummyRepoRepositorySuccess{},
		},
		{
			name:           "TestCreateRepo,::RepoRepositoryError",
			repo:           repo,
			wantErr:        true,
			resultErr:      errors.New(""),
			repoRepository: &interfaces.DummyRepoRepositoryFail{},
		},
	}

	for _, tt := range testList {
		t.Run(tt.name, func(t *testing.T) {
			service := NewRepoService(conf, tt.repoRepository, nil)
			_, err := service.CreateRepo(tt.repo)
			if err := tests.CompareError(tt.wantErr, err, tt.resultErr); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestUpdateRepo(t *testing.T) {
	conf := config.ReadDummy()
	id := uuid.New()
	repo := &models.Repo{}

	testList := []struct {
		name           string
		id             uuid.UUID
		repo           *models.Repo
		wantErr        bool
		resultErr      error
		repoRepository interfaces.RepoRepository
	}{
		{
			name:           "TestUpdateRepo::BestCase",
			id:             id,
			repo:           repo,
			wantErr:        false,
			resultErr:      nil,
			repoRepository: &interfaces.DummyRepoRepositorySuccess{},
		},
		{
			name:           "TestUpdateRepo,::RepoRepositoryError",
			id:             id,
			repo:           repo,
			wantErr:        true,
			resultErr:      errors.New(""),
			repoRepository: &interfaces.DummyRepoRepositoryFail{},
		},
	}

	for _, tt := range testList {
		t.Run(tt.name, func(t *testing.T) {
			service := NewRepoService(conf, tt.repoRepository, nil)
			err := service.UpdateRepo(tt.id, tt.repo)
			if err := tests.CompareError(tt.wantErr, err, tt.resultErr); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestDeleteRepo(t *testing.T) {
	conf := config.ReadDummy()
	id := uuid.New()

	testList := []struct {
		name           string
		id             uuid.UUID
		wantErr        bool
		resultErr      error
		repoRepository interfaces.RepoRepository
	}{
		{
			name:           "TestDeleteRepo::BestCase",
			id:             id,
			wantErr:        false,
			resultErr:      nil,
			repoRepository: &interfaces.DummyRepoRepositorySuccess{},
		},
		{
			name:           "TestDeleteRepo,::RepoRepositoryError",
			id:             id,
			wantErr:        true,
			resultErr:      errors.New(""),
			repoRepository: &interfaces.DummyRepoRepositoryFail{},
		},
	}

	for _, tt := range testList {
		t.Run(tt.name, func(t *testing.T) {
			service := NewRepoService(conf, tt.repoRepository, nil)
			err := service.DeleteRepo(tt.id)
			if err := tests.CompareError(tt.wantErr, err, tt.resultErr); err != nil {
				t.Error(err)
			}
		})
	}
}
