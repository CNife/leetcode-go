package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortIntSlices(t *testing.T) {
	assert.Equal(t,
		[][]int{
			nil,
			{7, 8},
			{4, 2, 1, 5, 3},
			{6, 1, 5, 3, 2},
		},
		IntSlices([][]int{
			{4, 2, 1, 5, 3},
			{6, 1, 5, 3, 2},
			{7, 8},
			nil,
		}),
	)
}

func TestSortIntSlicesDeep(t *testing.T) {
	assert.Equal(t,
		[][]int{
			nil,
			{7, 8},
			{1, 2, 3, 4, 5},
			{1, 2, 3, 5, 6},
		},
		IntSlicesDeep(
			[][]int{
				{4, 2, 1, 5, 3},
				{6, 1, 5, 3, 2},
				{7, 8},
				nil,
			},
			Ints,
		),
	)
}

func TestSortStringSlices(t *testing.T) {
	assert.Equal(t,
		[][]string{
			nil,
			{"bc", "ab", "ba", "b", "abc", "c", "a"},
		},
		StringSlices([][]string{
			{"bc", "ab", "ba", "b", "abc", "c", "a"},
			nil,
		}),
	)
}

func TestSortStringSlicesDeep(t *testing.T) {
	assert.Equal(t,
		[][]string{
			nil,
			{"a", "ab", "abc", "b", "ba", "bc", "c"},
		},
		StringSlicesDeep(
			[][]string{
				{"bc", "ab", "ba", "b", "abc", "c", "a"},
				nil,
			},
			Strings,
		),
	)
}
