package services

import (
	"errors"
	"testing"

	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/internal/constants"
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

func TestMatchFindingRule(t *testing.T) {
	testList := []struct {
		name        string
		findingRule []config.FindingRule
		word        string
	}{
		{
			name: "TestMatchFindingRule::Prefix",
			findingRule: []config.FindingRule{
				{
					Match: constants.FindingRuleMatchPrefix,
					Type:  "God",
				},
			},
			word: "GodEmperor",
		},
		{
			name: "TestMatchFindingRule::Suffix",
			findingRule: []config.FindingRule{
				{
					Match: constants.FindingRuleMatchSuffix,
					Type:  "Emperor",
				},
			},
			word: "GodEmperor",
		},
		{
			name: "TestMatchFindingRule::Whole",
			findingRule: []config.FindingRule{
				{
					Match: constants.FindingRuleMatchWhole,
					Type:  "GodEmperor",
				},
			},
			word: "GodEmperor",
		},
		{
			name: "TestMatchFindingRule::Partial",
			findingRule: []config.FindingRule{
				{
					Match: constants.FindingRuleMatchPartial,
					Type:  "dEmpe",
				},
			},
			word: "GodEmperor",
		},
	}

	for _, tt := range testList {
		t.Run(tt.name, func(t *testing.T) {
			if result := matchFindingRule(tt.findingRule, tt.word); result == nil {
				t.Error(tests.ErrNoResult)
			}
		})
	}
}
