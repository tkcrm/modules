package dbutils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tkcrm/modules/pkg/db/dbutils"
	"github.com/tkcrm/modules/pkg/utils"
)

func Test_Pagination(t *testing.T) {
	tests := []struct {
		page           *uint64
		pageSize       *uint64
		expectedLimit  uint64
		expectedOffset uint64
		opts           []dbutils.PaginationOption
	}{
		{
			page:           utils.Pointer(uint64(1)),
			pageSize:       utils.Pointer(uint64(1)),
			expectedLimit:  1,
			expectedOffset: 0,
		},
		{
			page:           utils.Pointer(uint64(0)),
			pageSize:       utils.Pointer(uint64(10)),
			expectedLimit:  10,
			expectedOffset: 0,
		},
		{
			page:           utils.Pointer(uint64(1)),
			pageSize:       utils.Pointer(uint64(1000)),
			expectedLimit:  1000,
			expectedOffset: 0,
			opts:           []dbutils.PaginationOption{dbutils.WithMaxLimit(1000)},
		},
		{
			page:           utils.Pointer(uint64(2)),
			pageSize:       utils.Pointer(uint64(1000)),
			expectedLimit:  1000,
			expectedOffset: 1000,
			opts:           []dbutils.PaginationOption{dbutils.WithMaxLimit(1000)},
		},
		{
			page:           nil,
			pageSize:       nil,
			expectedLimit:  100,
			expectedOffset: 0,
			opts:           []dbutils.PaginationOption{dbutils.WithMaxLimit(25000)},
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			limit, offset, err := dbutils.Pagination(tc.page, tc.pageSize, tc.opts...)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedLimit, limit)
			assert.Equal(t, tc.expectedOffset, offset)
		})
	}
}
