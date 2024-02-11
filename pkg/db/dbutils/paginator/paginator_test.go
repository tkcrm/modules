package paginator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tkcrm/modules/pkg/db/dbutils/paginator"
)

func TestPaginator(t *testing.T) {
	tests := []struct {
		page, pageSize, totalItems int
		expexted                   []*paginator.DrawPagesItem
	}{
		{
			page:       1,
			pageSize:   3,
			totalItems: 20,
			expexted: []*paginator.DrawPagesItem{
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
			expexted: []*paginator.DrawPagesItem{
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
			expexted: []*paginator.DrawPagesItem{
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
			page:       7,
			pageSize:   3,
			totalItems: 20,
			expexted: []*paginator.DrawPagesItem{
				{
					PageNumber: 5,
				},
				{
					PageNumber: 6,
				},
				{
					PageNumber: 7,
					IsActive:   true,
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			p := paginator.New(tc.page, tc.pageSize, tc.totalItems, paginator.WithMaxVisibleItems(3))
			res := p.DrawPages()
			assert.Equal(t, tc.expexted, res)
		})
	}
}
