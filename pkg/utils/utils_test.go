package utils_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tkcrm/modules/pkg/utils"
)

func Test_FilterArray(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			in:  []int{1, 2, 3, 4},
			out: []int{1, 2},
		},
	}

	t.Run("filter ints", func(t *testing.T) {
		for _, tc := range tests {
			res := utils.FilterArray(tc.in, func(v int) bool {
				return v < 3
			})

			assert.Equal(t, tc.out, res)
		}
	})
}

func Test_FilterValues(t *testing.T) {
	tests := []struct {
		in       []int
		filter   []int
		expected []int
	}{
		{
			in:       []int{1, 2, 3, 4},
			filter:   []int{1, 2},
			expected: []int{3, 4},
		},
	}

	t.Run("filter ints", func(t *testing.T) {
		for _, tc := range tests {
			res := utils.FilterValues(tc.in, tc.filter)

			assert.Equal(t, tc.expected, res)
		}
	})
}

func Test_FindInArray(t *testing.T) {
	type inItem struct {
		key int
	}
	tests := []struct {
		in       []inItem
		expected bool
	}{
		{
			in: []inItem{
				{
					key: 1,
				},
				{
					key: 2,
				},
			},
			expected: false,
		},
		{
			in: []inItem{
				{
					key: 5,
				},
				{
					key: 4,
				},
			},
			expected: false,
		},
		{
			in: []inItem{
				{
					key: 6,
				},
				{
					key: 4,
				},
			},
			expected: true,
		},
		{
			in: []inItem{
				{
					key: 6,
				},
				{
					key: 7,
				},
			},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			_, ok := utils.FindInArray(tc.in, func(v inItem) bool {
				return v.key > 5
			})

			if ok != tc.expected {
				t.Fatal("fail")
			}
		})
	}
}

func TestGenerateRandomNumber(t *testing.T) {
	tests := []struct {
		length uint
		min    int
		max    int
	}{
		{
			length: 1,
			min:    1,
			max:    9,
		},
		{
			length: 2,
			min:    10,
			max:    99,
		},
		{
			length: 6,
			min:    100000,
			max:    999999,
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			res := utils.GenerateRandomNumber(tc.length)
			if res < tc.min {
				t.Fatal("less")
			}
			if res > tc.max {
				t.Fatal("greater")
			}
			fmt.Println(res)
		})
	}
}
