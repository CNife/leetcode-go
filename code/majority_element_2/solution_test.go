package majority_element_2

import (
	"testing"

	"github.com/CNife/leetcode-go/util"
	"github.com/stretchr/testify/assert"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums, want []int
	}{
		{
			nums: []int{3, 2, 3},
			want: []int{3},
		},
		{
			nums: []int{1},
			want: []int{1},
		},
		{
			nums: []int{1, 1, 1, 3, 3, 2, 2, 2},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, util.SortInts(tt.want),
			util.SortInts(MajorityElement(tt.nums)))
	}
}
