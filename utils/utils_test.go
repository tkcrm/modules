package utils_test

import (
	"testing"

	"github.com/tkcrm/modules/utils"
)

func Test_ExistInArray(t *testing.T) {
	tt := []struct {
		arr      []string
		value    string
		expected bool
	}{
		{
			arr:      []string{"A", "B", "C"},
			value:    "C",
			expected: true,
		},
		{
			arr:      []string{"A", "B", "C"},
			value:    "D",
			expected: false,
		},
	}

	for index, ts := range tt {
		if utils.ExistInArray(ts.arr, ts.value) != ts.expected {
			t.Fatalf("%d test error:", index)
		}
	}
}
