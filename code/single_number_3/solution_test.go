package single_number_3

import (
	"testing"

	"github.com/CNife/leetcode-go/util"
	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		nums, want []int
	}{
		{
			nums: []int{1, 2, 1, 3, 2, 5},
			want: []int{3, 5},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, util.SortInts(tt.want),
			util.SortInts(SingleNumber(tt.nums)))
	}
}
