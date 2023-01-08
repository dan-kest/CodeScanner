package tables

import (
	"time"

	"github.com/google/uuid"
)

type ScanResult struct {
	TransactionID uuid.UUID `gorm:"type:uuid;primaryKey" json:"transaction_id"`
	Result        string    `gorm:"type:text" json:"result"`
	CreatedAt     time.Time `gorm:"type:timestamp(3)" json:"created_at"`
}

func (m *ScanResult) TableName() string {
	return "scan_result"
}
