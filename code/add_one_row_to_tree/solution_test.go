package add_one_row_to_tree

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestAddOneRow(t *testing.T) {
	tests := []struct {
		root       *types.TreeNode
		val, depth int
		want       *types.TreeNode
	}{
		{
			root:  types.NewTree(4, 2, 6, 3, 1, 5),
			val:   1,
			depth: 2,
			want:  types.NewTree(4, 1, 1, 2, -1, -1, 6, 3, 1, 5),
		},
		{
			root:  types.NewTree(4, 2, -1, 3, 1),
			val:   1,
			depth: 3,
			want:  types.NewTree(4, 2, -1, 1, 1, 3, -1, -1, 1),
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, AddOneRow(tt.root, tt.val, tt.depth))
	}
}
