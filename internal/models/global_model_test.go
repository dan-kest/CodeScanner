package models

import (
	"testing"

	"github.com/dan-kest/cscanner/config"
	"github.com/dan-kest/cscanner/pkg/tests"
)

func TestIsSortDesc(t *testing.T) {
	testList := []struct {
		name   string
		paging Paging
		result bool
	}{
		{
			name: "TestIsSortDesc::SortOrderEmpty",
			paging: Paging{
				SortOrder: "",
			},
			result: false,
		},
		{
			name: "TestIsSortDesc::SortOrderRandomWord",
			paging: Paging{
				SortOrder: "asdf",
			},
			result: false,
		},
		{
			name: "TestIsSortDesc::SortOrderMixedCase",
			paging: Paging{
				SortOrder: "dEsC",
			},
			result: true,
		},
	}

	for _, tt := range testList {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.paging.IsSortDesc()
			if err := tests.CompareField(result, tt.result, nil); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestSetFallback(t *testing.T) {
	conf := config.ReadDummy()
	conf.App.Paging.MaxItemPerPage = 20

	testList := []struct {
		name   string
		paging Paging
		result Paging
	}{
		{
			name: "TestSetFallback::Page=0",
			paging: Paging{
				Page: 0,
			},
			result: Paging{
				Page: 1,
			},
		},
		{
			name: "TestSetFallback::Page<0",
			paging: Paging{
				Page: -4,
			},
			result: Paging{
				Page: 1,
			},
		},
		{
			name: "TestSetFallback::ItemPerPage<0",
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
			name: "TestSetFallback::ItemPerPage>Max",
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
			name: "TestSetFallback::MixedConfig",
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
			name: "TestSetFallback::MixedConfigReverse",
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

	for _, tt := range testList {
		t.Run(tt.name, func(t *testing.T) {
			tt.paging.SetFallback(conf.App.Paging)
			if err := tests.CompareField(tt.paging.Page, tt.result.Page, "Page"); err != nil {
				t.Error(err)
			}
			if err := tests.CompareField(tt.paging.ItemPerPage, tt.result.ItemPerPage, "ItemPerPage"); err != nil {
				t.Error(err)
			}
		})
	}
}
