package binary_tree_zigzag_level_order_traversal

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestZigzagLevelOrder(t *testing.T) {
	tests := []struct {
		root *types.TreeNode
		want [][]int
	}{
		{
			root: types.NewTree(3, 9, 20, -1, -1, 15, 7),
			want: [][]int{{3}, {20, 9}, {15, 7}},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, ZigzagLevelOrder(tt.root))
	}
}
