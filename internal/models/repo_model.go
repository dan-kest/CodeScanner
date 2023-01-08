package models

import (
	"github.com/dan-kest/cscanner/pkg/null"
	"github.com/google/uuid"
)

type RepoPagination struct {
	Page        int
	ItemPerPage int
	TotalCount  int
	ItemList    []*Repo
}

type Repo struct {
	ID         uuid.UUID
	Name       null.String
	URL        null.String
	ScanStatus ScanStatus
	Findings   []*Finding `json:"findings"`
	Timestamp  string
}

type ScanStatus string

const (
	ScanStatusQueued     ScanStatus = "Queued"
	ScanStatusInProgress ScanStatus = "In Progress"
	ScanStatusSuccess    ScanStatus = "Success"
	ScanStatusFailure    ScanStatus = "Failure"
)

type Finding struct {
	Type     string          `json:"type"`
	RuleID   string          `json:"ruleId"`
	Location FindingLocation `json:"location"`
	Metadata FindingMetadata `json:"metadata"`
}

type FindingLocation struct {
	Path      string                  `json:"path"`
	Positions FindingLocationPosition `json:"positions"`
}

type FindingLocationPosition struct {
	Begin FindingLocationPositionBegin `json:"begin"`
}

type FindingLocationPositionBegin struct {
	Line int `json:"line"`
}

type FindingMetadata struct {
	Description string `json:"description"`
	Severity    string `json:"severity"`
}
