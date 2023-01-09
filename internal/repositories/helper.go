package repositories

import (
	"github.com/dan-kest/cscanner/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Build pagination query for GORM
func buildPaging(tx *gorm.DB, paging *models.Paging) *gorm.DB {
	if paging.ItemPerPage > 0 {
		tx = tx.Limit(paging.ItemPerPage)
		tx = tx.Offset((paging.Page - 1) * paging.ItemPerPage)
	}

	if paging.SortBy != "" {
		tx.Order(clause.OrderByColumn{
			Column: clause.Column{
				Name: paging.SortBy,
			},
			Desc: paging.IsSortDesc(),
		})
	}

	return tx
}
