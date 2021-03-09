package sum_root_to_leaf_numbers

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestSumNumbers(t *testing.T) {
	tests := []struct {
		root *types.TreeNode
		want int
	}{
		{
			root: types.NewTree(1, 2, 3),
			want: 25,
		},
		{
			root: types.NewTree(4, 9, 0, 5, 1),
			want: 1026,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, SumNumbers(tt.root))
	}
}
