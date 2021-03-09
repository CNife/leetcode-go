package find_all_continuous_sequences

import (
	"testing"

	"github.com/CNife/leetcode-go/util"
	"github.com/stretchr/testify/assert"
)

func TestFindContinuousSequence(t *testing.T) {
	tests := []struct {
		target int
		want   [][]int
	}{
		{
			target: 9,
			want: [][]int{
				{2, 3, 4},
				{4, 5},
			},
		},
		{
			target: 15,
			want: [][]int{
				{1, 2, 3, 4, 5},
				{4, 5, 6},
				{7, 8},
			},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, util.SortIntSlices(tt.want),
			util.SortIntSlices(FindContinuousSequence(tt.target)))
	}
}
