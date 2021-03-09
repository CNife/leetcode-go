package intersection_of_two_arrays

import (
	"testing"

	"github.com/CNife/leetcode-go/util"
	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		arr1, arr2, want []int
	}{
		{
			arr1: []int{1, 2, 2, 1},
			arr2: []int{2, 2},
			want: []int{2},
		},
		{
			arr1: []int{4, 9, 5},
			arr2: []int{9, 4, 9, 8, 4},
			want: []int{9, 4},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, util.SortInts(tt.want),
			util.SortInts(Intersection(tt.arr1, tt.arr2)))
	}
}
