package constants

const (
	ErrorNotFoundSuffix = "not found"
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
	RabbitMQWaitingPrompt     = " [*] Waiting for messages. To exit press CTRL+C"
	RabbitMQErrorOpenChannel  = "Failed to open a channel"
	RabbitMQErrorDeclareQueue = "Failed to declare a queue"
	RabbitMQErrorPublish      = "Failed to publish a message"
	RabbitMQErrorConsume      = "Failed to register a consumer"
)

const (
	MIMETypeApplicationJSON = "application/json"
)
