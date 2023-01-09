package models

import (
	"github.com/dan-kest/cscanner/internal/constants"
	"github.com/google/uuid"
)

type Task struct {
	RepositoryIDStr string               `json:"repository_id"`
	RepositoryID    uuid.UUID            `json:"-"`
	ScanIDStr       string               `json:"scan_id"`
	ScanID          uuid.UUID            `json:"-"`
	URL             string               `json:"url"`
	Timestamp       string               `json:"timestamp"`
	Status          constants.ScanStatus `json:"-"`
}
