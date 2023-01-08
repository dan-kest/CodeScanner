package tables

import (
	"time"
)

type ScanHistory struct {
	ID           uint        `gorm:"primaryKey" json:"id"`
	RepositoryID uint        `gorm:"index" json:"repository_id"`
	Status       string      `json:"status"`
	CreatedAt    time.Time   `json:"created_at"`
	Repository   *Repository `gorm:"foreignKey:RepositoryID" json:"repository"`
}

func (m *ScanHistory) TableName() string {
	return "scan_history"
}
