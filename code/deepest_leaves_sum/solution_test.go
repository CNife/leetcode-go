package deepest_leaves_sum

import (
	"testing"

	. "github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestDeepestLeaveSum(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{
			root: NewTree(1, 2, 3, 4, 5, -1, 6, 7, -1, -1, -1, -1, 8),
			want: 15,
		},
		{
			root: NewTree(6, 7, 8, 2, 7, 1, 3, 9, -1, 1, 4, -1, -1, -1, 5),
			want: 19,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, DeepestLeaveSum(tt.root))
	}
}
