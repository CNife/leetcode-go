package increasing_order_search_tree

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestIncreasingBST(t *testing.T) {
	tests := []struct {
		root, want *types.TreeNode
	}{
		{
			root: types.NewTree(5, 3, 6, 2, 4, -1, 8, 1, -1, -1, -1, 7, 9),
			want: types.NewTree(1, -1, 2, -1, 3, -1, 4, -1, 5, -1, 6, -1, 7, -1, 8, -1, 9),
		},
		{
			root: types.NewTree(5, 1, 7),
			want: types.NewTree(1, -1, 5, -1, 7),
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, IncreasingBST(tt.root))
	}
}
