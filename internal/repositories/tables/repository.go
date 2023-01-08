package tables

import (
	"time"

	"github.com/dan-kest/cscanner/pkg/null"
	"gorm.io/gorm"
)

type Repository struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      null.String    `json:"name"`
	URL       null.String    `json:"url"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (m *Repository) TableName() string {
	return "repository"
}
