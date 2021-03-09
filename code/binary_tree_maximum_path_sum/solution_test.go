package binary_tree_maximum_path_sum

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		tree *types.TreeNode
		want int
	}{
		{
			tree: &types.TreeNode{
				Val:   1,
				Left:  &types.TreeNode{Val: 2},
				Right: &types.TreeNode{Val: 3},
			},
			want: 6,
		},
		{
			tree: &types.TreeNode{
				Val:  -10,
				Left: &types.TreeNode{Val: 9},
				Right: &types.TreeNode{
					Val: 20,
					Left: &types.TreeNode{
						Val: 15,
					},
					Right: &types.TreeNode{
						Val: 7,
					},
				},
			},
			want: 42,
		},
		{
			tree: &types.TreeNode{Val: -3},
			want: -3,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, MaxPathSum(tt.tree))
	}
}
