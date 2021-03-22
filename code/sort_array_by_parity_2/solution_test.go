package sort_array_by_parity_2

import (
	"testing"

	"github.com/CNife/leetcode-go/util/sort"
	"github.com/stretchr/testify/assert"
)

func TestSortArrayByParity(t *testing.T) {
	arrays := [][]int{
		{4, 2, 5, 7},
	}
	for _, array := range arrays {
		got := SortArrayByParity(array)
		for i, num := range got {
			if (i^num)&1 != 0 {
				t.Errorf("parity error at %v of %v", i, got)
			}
		}
		assert.Equal(t, sort.Ints(array), sort.Ints(got))
	}
}
