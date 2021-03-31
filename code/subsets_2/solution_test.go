package subsets_2

import (
	"testing"

	"github.com/CNife/leetcode-go/util/sort"
	"github.com/stretchr/testify/assert"
)

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{1, 2, 2},
			want: [][]int{
				{},
				{1},
				{1, 2},
				{1, 2, 2},
				{2},
				{2, 2},
			},
		},
		{
			nums: []int{0},
			want: [][]int{
				{},
				{0},
			},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, sort.IntSlicesDeep(tt.want, sort.Ints),
			sort.IntSlicesDeep(SubsetsWithDup(tt.nums), sort.Ints))
	}
}
