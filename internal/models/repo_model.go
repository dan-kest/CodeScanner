package models

import (
	"github.com/dan-kest/cscanner/pkg/null"
)

type Repo struct {
	ID         int
	Name       null.String
	URL        null.String
	ScanStatus ScanStatus
	Findings   []*Finding `json:"findings"`
	Timestamp  string
}

type ScanStatus string

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
