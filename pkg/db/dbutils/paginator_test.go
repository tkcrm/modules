package dbutils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tkcrm/modules/pkg/db/dbutils"
)

func TestPaginator(t *testing.T) {
	tests := []struct {
		page, pageSize, totalItems int
		expexted                   []dbutils.DrawPagesItem
	}{
		{
			page:       1,
			pageSize:   3,
			totalItems: 20,
			expexted: []dbutils.DrawPagesItem{
				{
					PageNumber: 1,
					IsActive:   true,
				},
				{
					PageNumber: 2,
				},
				{
					PageNumber: 3,
				},
			},
		},
		{
			page:       2,
			pageSize:   3,
			totalItems: 20,
			expexted: []dbutils.DrawPagesItem{
				{
					PageNumber: 1,
				},
				{
					PageNumber: 2,
					IsActive:   true,
				},
				{
					PageNumber: 3,
				},
			},
		},
		{
			page:       3,
			pageSize:   3,
			totalItems: 20,
			expexted: []dbutils.DrawPagesItem{
				{
					PageNumber: 2,
				},
				{
					PageNumber: 3,
					IsActive:   true,
				},
				{
					PageNumber: 4,
				},
			},
		},
		{
			page:       6,
			pageSize:   3,
			totalItems: 20,
			expexted: []dbutils.DrawPagesItem{
				{
					PageNumber: 4,
				},
				{
					PageNumber: 5,
				},
				{
					PageNumber: 6,
					IsActive:   true,
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			p := dbutils.New(tc.page, tc.pageSize, tc.totalItems)
			res := p.DrawPages()
			assert.Equal(t, tc.expexted, res)
		})
	}
}
