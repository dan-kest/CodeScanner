package models

import (
	"time"

	"github.com/dan-kest/cscanner/internal/constants"
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
	ID         uuid.UUID            `json:"-"`
	Name       null.String          `json:"-"`
	URL        null.String          `json:"-"`
	ScanStatus constants.ScanStatus `json:"-"`
	Findings   []*Finding           `json:"findings"`
	Timestamp  *time.Time           `json:"-"`
}

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
