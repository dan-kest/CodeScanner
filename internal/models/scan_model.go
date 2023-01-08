package models

import "github.com/google/uuid"

type Task struct {
	RepositoryIDStr string `json:"repository_id"`
	RepositoryID    uuid.UUID
	ScanIDStr       string `json:"scan_id"`
	ScanID          uuid.UUID
	URL             string `json:"url"`
	Timestamp       string `json:"timestamp"`
	Status          ScanStatus
}
