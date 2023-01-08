package tables

import (
	"time"

	"github.com/dan-kest/cscanner/pkg/null"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	ID              uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v1()" json:"id"`
	Name            null.String    `gorm:"type:varchar(100);not null" json:"name"`
	URL             null.String    `gorm:"type:varchar(255);not null" json:"url"`
	CreatedAt       time.Time      `gorm:"type:timestamp(3)" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"type:timestamp(3)" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"type:timestamp(3);index" json:"deleted_at"`
	ScanHistoryList []*ScanHistory `gorm:"foreignKey:RepositoryID" json:"scan_history_list"`
}

func (m *Repository) TableName() string {
	return "repository"
}
