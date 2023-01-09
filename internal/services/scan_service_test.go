package services

import (
	"errors"
	"testing"

	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/internal/interfaces"
	"github.com/dan-kest/cscanner/internal/models"
	"github.com/dan-kest/cscanner/pkg/tests"
)

func TestRunTask(t *testing.T) {
	conf := config.ReadDummy()

	testList := []struct {
		name           string
		task           *models.Task
		wantErr        bool
		resultErr      error
		scanRepository interfaces.ScanRepository
	}{
		{
			name:           "TestRunTask::BestCase",
			task:           &models.Task{},
			wantErr:        false,
			resultErr:      nil,
			scanRepository: &interfaces.DummyScanRepositorySuccess{},
		},
		{
			name:           "TestRunTask::ScanRepositoryError",
			task:           &models.Task{},
			wantErr:        true,
			resultErr:      errors.New(""),
			scanRepository: &interfaces.DummyScanRepositoryFail{},
		},
	}

	for _, tt := range testList {
		t.Run(tt.name, func(t *testing.T) {
			service := NewScanService(conf, tt.scanRepository)
			err := service.RunTask(tt.task)
			if err := tests.CompareError(tt.wantErr, err, tt.resultErr); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestRunErrorTask(t *testing.T) {
	conf := config.ReadDummy()

	testList := []struct {
		name           string
		body           []byte
		err            error
		wantErr        bool
		resultErr      error
		scanRepository interfaces.ScanRepository
	}{
		{
			name:           "TestRunErrorTask::BestCase",
			body:           []byte{},
			err:            errors.New(""),
			wantErr:        false,
			resultErr:      nil,
			scanRepository: &interfaces.DummyScanRepositorySuccess{},
		},
		{
			name:           "TestRunErrorTask::ScanRepositoryError",
			body:           []byte{},
			err:            errors.New(""),
			wantErr:        true,
			resultErr:      errors.New(""),
			scanRepository: &interfaces.DummyScanRepositoryFail{},
		},
	}

	for _, tt := range testList {
		t.Run(tt.name, func(t *testing.T) {
			service := NewScanService(conf, tt.scanRepository)
			err := service.RunErrorTask(tt.body, tt.err)
			if err := tests.CompareError(tt.wantErr, err, tt.resultErr); err != nil {
				t.Error(err)
			}
		})
	}
}
