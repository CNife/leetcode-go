package count_complete_tree_nodes

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestCountNodes(t *testing.T) {
	tests := []struct {
		root *types.TreeNode
		want int
	}{
		{
			root: nil,
			want: 0,
		},
		{

			root: types.NewTree(1, 2, 3, 4, 5, 6),
			want: 6,
		},
		{
			root: types.NewTree(1, 2, 3),
			want: 3,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, CountNodes(tt.root))
	}
}
