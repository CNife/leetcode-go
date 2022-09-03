package longest_univalue_path

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestLongestUnivaluePath(t *testing.T) {
	tests := []struct {
		root *types.TreeNode
		want int
	}{
		{
			root: types.NewTree(5, 4, 5, 1, 1, 5),
			want: 2,
		},
		{
			root: types.NewTree(1, 4, 5, 4, 4, 5),
			want: 2,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, LongestUnivaluePath(tt.root))
	}
}
