package tables

import (
	"time"

	"github.com/google/uuid"
)

type ScanHistory struct {
	ID           int         `gorm:"type:int8;primaryKey;autoIncrement" json:"id"`
	RepositoryID uuid.UUID   `gorm:"type:uuid;not null;index" json:"repository_id"`
	ScanID       uuid.UUID   `gorm:"type:uuid;not null" json:"transaction_id"`
	Status       string      `gorm:"type:varchar(11);not null" json:"status"`
	CreatedAt    time.Time   `gorm:"type:timestamp(6)" json:"created_at"`
	Repository   *Repository `gorm:"foreignKey:RepositoryID" json:"repository"`
	ScanResult   *ScanResult `gorm:"foreignKey:ScanID" json:"scan_result"`
}

func (m *ScanHistory) TableName() string {
	return "scan_history"
}
