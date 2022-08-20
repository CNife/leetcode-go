package maximum_binary_tree

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestConstructMaximumBinaryTree(t *testing.T) {
	tests := []struct {
		nums []int
		want *types.TreeNode
	}{
		{
			nums: []int{3, 2, 1, 6, 0, 5},
			want: types.NewTree(6, 3, 5, -1, 2, 0, -1, -1, 1),
		},
		{
			nums: []int{3, 2, 1},
			want: types.NewTree(3, -1, 2, -1, 1),
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, ConstructMaximumBinaryTree(tt.nums))
	}
}
