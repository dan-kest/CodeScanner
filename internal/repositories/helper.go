package repositories

import (
	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func buildPaging(tx *gorm.DB, conf *config.Config, paging *models.Paging) *gorm.DB {
	if paging.ItemPerPage > 0 && paging.Page > 0 {
		if paging.ItemPerPage > conf.App.Paging.MaxItemPerPage {
			paging.ItemPerPage = conf.App.Paging.MaxItemPerPage
		}
		tx = tx.Limit(paging.ItemPerPage)

		if paging.Page == 0 {
			paging.Page = 1
		}
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
