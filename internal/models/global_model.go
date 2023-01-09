package models

import (
	"strings"

	"github.com/dan-kest/cscanner/config"
)

type Paging struct {
	Page        int
	ItemPerPage int
	SortBy      string
	SortOrder   string
}

// Return true if sort by DESC
func (m *Paging) IsSortDesc() bool {
	return strings.ToLower(m.SortOrder) == "desc"
}

// Set fallback value if value is invalid
func (m *Paging) SetFallback(conf *config.Paging) {
	if m.Page <= 0 {
		m.Page = 1
	}
	if m.ItemPerPage < 0 {
		m.ItemPerPage = 0
	} else if m.ItemPerPage > conf.MaxItemPerPage {
		m.ItemPerPage = conf.MaxItemPerPage
	}
}
