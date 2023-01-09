package models

import (
	"testing"

	"github.com/dan-kest/cscanner/config"
)

var (
	conf = config.Config{
		App: &config.App{
			Paging: &config.Paging{
				MaxItemPerPage: 20,
			},
		},
	}
)

func TestIsSortDesc(t *testing.T) {
	tests := []struct {
		name   string
		paging Paging
		result bool
	}{
		{
			name: "TestIsSortDesc__SortOrderEmpty",
			paging: Paging{
				SortOrder: "",
			},
			result: false,
		},
		{
			name: "TestIsSortDesc__SortOrderRandomWord",
			paging: Paging{
				SortOrder: "asdf",
			},
			result: false,
		},
		{
			name: "TestIsSortDesc__SortOrderMixedCase",
			paging: Paging{
				SortOrder: "dEsC",
			},
			result: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.paging.IsSortDesc()
			if result != tt.result {
				t.Errorf("got '%t', want '%t'", result, tt.result)
			}
		})
	}
}

func TestSetFallback(t *testing.T) {
	tests := []struct {
		name   string
		paging Paging
		result Paging
	}{
		{
			name: "TestSetFallback__Page=0",
			paging: Paging{
				Page: 0,
			},
			result: Paging{
				Page: 1,
			},
		},
		{
			name: "TestSetFallback__Page<0",
			paging: Paging{
				Page: -4,
			},
			result: Paging{
				Page: 1,
			},
		},
		{
			name: "TestSetFallback__ItemPerPage<0",
			paging: Paging{
				Page:        1,
				ItemPerPage: -2,
			},
			result: Paging{
				Page:        1,
				ItemPerPage: 0,
			},
		},
		{
			name: "TestSetFallback__ItemPerPage>Max",
			paging: Paging{
				Page:        1,
				ItemPerPage: 50,
			},
			result: Paging{
				Page:        1,
				ItemPerPage: conf.App.Paging.MaxItemPerPage,
			},
		},
		{
			name: "TestSetFallback__MixedConfig",
			paging: Paging{
				Page:        3,
				ItemPerPage: 50,
			},
			result: Paging{
				Page:        3,
				ItemPerPage: conf.App.Paging.MaxItemPerPage,
			},
		},
		{
			name: "TestSetFallback__MixedConfigReverse",
			paging: Paging{
				Page:        -2,
				ItemPerPage: 5,
			},
			result: Paging{
				Page:        1,
				ItemPerPage: 5,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.paging.SetFallback(conf.App.Paging)
			if tt.paging.Page != tt.result.Page {
				t.Errorf("Page: got '%d', want '%d'", tt.paging.Page, tt.result.Page)
			}
			if tt.paging.ItemPerPage != tt.result.ItemPerPage {
				t.Errorf("ItemPerPage: got '%d', want '%d'", tt.paging.ItemPerPage, tt.result.ItemPerPage)
			}
		})
	}
}
