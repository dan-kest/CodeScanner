package constants

const (
	MessageRabbitMQWaitingPrompt = " [*] Waiting for messages. To exit press CTRL+C"
)

const (
	ErrorNotFoundSuffix       = "not found"
	ErrorTaskFailed           = "Error task failed"
	ErrorRabbitMQOpenChannel  = "Failed to open a channel"
	ErrorRabbitMQDeclareQueue = "Failed to declare a queue"
	ErrorRabbitMQAcknowledge  = "Failed to acknowledge a message"
	ErrorRabbitMQPublish      = "Failed to publish a message"
	ErrorRabbitMQConsume      = "Failed to register a consumer"
)

type ScanStatus string

const (
	ScanStatusQueued     ScanStatus = "Queued"
	ScanStatusInProgress ScanStatus = "In Progress"
	ScanStatusSuccess    ScanStatus = "Success"
	ScanStatusFailure    ScanStatus = "Failure"
)

const (
	FindingSASTType     = "sast"
	FindingSASTRuleID   = "G402"
	FindingExposedDesc  = "Exposed sensitive information."
	FindingSeverityHigh = "HIGH"
)

const (
	ResponseStatusOK    = "OK"
	ResponseStatusError = "ERROR"
)

const (
	MIMETypeApplicationJSON = "application/json"
)

const (
	ScanWordDelimiter = " "
)

const (
	FindingRuleMatchPrefix  = "prefix"
	FindingRuleMatchSuffix  = "suffix"
	FindingRuleMatchWhole   = "whole"
	FindingRuleMatchPartial = "partial"
)
