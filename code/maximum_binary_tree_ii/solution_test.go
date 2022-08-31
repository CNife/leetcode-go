package maximum_binary_tree_ii

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestInsertIntoMaxTree(t *testing.T) {
	tests := []struct {
		root *types.TreeNode
		val  int
		want *types.TreeNode
	}{
		{
			root: types.NewTree(4, 1, 3, -1, -1, 2),
			val:  5,
			want: types.NewTree(5, 4, -1, 1, 3, -1, -1, 2),
		},
		{
			root: types.NewTree(5, 2, 4, -1, 1),
			val:  3,
			want: types.NewTree(5, 2, 4, -1, 1, -1, 3),
		},
		{
			root: types.NewTree(5, 2, 3, -1, 1),
			val:  4,
			want: types.NewTree(5, 2, 4, -1, 1, 3),
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, InsertIntoMaxTree(tt.root, tt.val))
	}
}
