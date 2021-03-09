package recover_a_tree_from_preorder_traversal

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestRecoverFromPreorder(t *testing.T) {
	tests := []struct {
		s    string
		want *types.TreeNode
	}{
		{
			s:    "1-2--3--4-5--6--7",
			want: types.NewTree(1, 2, 5, 3, 4, 6, 7),
		},
		{
			s:    "1-2--3---4-5--6---7",
			want: types.NewTree(1, 2, 5, 3, -1, 6, -1, 4, -1, 7),
		},
		{
			s:    "1-401--349---90--88",
			want: types.NewTree(1, 401, -1, 349, 88, 90),
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, RecoverFromPreorder(tt.s))
	}
}
