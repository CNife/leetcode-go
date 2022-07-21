package binary_tree_pruning

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestPruneTree(t *testing.T) {
	tests := []struct {
		root, want *types.TreeNode
	}{
		{
			root: types.NewTree(1, -1, 0, 0, 1),
			want: types.NewTree(1, -1, 0, -1, 1),
		},
		{
			root: types.NewTree(1, 0, 1, 0, 0, 0, 1),
			want: types.NewTree(1, -1, 1, -1, 1),
		},
		{
			root: types.NewTree(1, 1, 0, 1, 1, 0, 1, 0),
			want: types.NewTree(1, 1, 0, 1, 1, -1, 1),
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, PruneTree(tt.root))
	}
}
