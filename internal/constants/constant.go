package constants

import "github.com/dan-kest/cscanner/internal/models"

const (
	ScanStatusQueued     models.ScanStatus = "Queued"
	ScanStatusInProgress models.ScanStatus = "In Progress"
	ScanStatusSuccess    models.ScanStatus = "Success"
	ScanStatusFailure    models.ScanStatus = "Failure"
)

const (
	FindingSASTType     = "sast"
	FindingSASTRuleID   = "G402"
	FindingExposedDesc  = "Exposed sensitive information."
	FindingSeverityHigh = "HIGH"
)
