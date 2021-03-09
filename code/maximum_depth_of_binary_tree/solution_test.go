package maximum_depth_of_binary_tree

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		root *types.TreeNode
		want int
	}{
		{
			root: types.NewTree(3, 9, 20, -1, -1, 15, 7),
			want: 3,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MaxDepth(tt.root))
	}
}
