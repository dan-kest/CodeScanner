package models

import "strings"

type Paging struct {
	Page        int
	ItemPerPage int
	SortBy      string
	SortOrder   string
}

func (m *Paging) IsSortDesc() bool {
	return strings.ToLower(m.SortOrder) == "desc"
}
