package unique_binary_search_trees_2

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestGenerateTrees(t *testing.T) {
	tests := []struct {
		n    int
		want []*types.TreeNode
	}{
		{
			n: 2,
			want: []*types.TreeNode{
				types.NewTree(1, -1, 2),
				types.NewTree(2, 1),
			},
		},
		{
			n: 1,
			want: []*types.TreeNode{
				types.NewTree(1),
			},
		},
		{
			n: 3,
			want: []*types.TreeNode{
				types.NewTree(1, -1, 3, 2),
				types.NewTree(3, 2, -1, 1),
				types.NewTree(3, 1, -1, -1, 2),
				types.NewTree(2, 1, 3),
				types.NewTree(1, -1, 2, -1, 3),
			},
		},
		{
			n:    0,
			want: nil,
		},
	}

	slice2Set := func(trees []*types.TreeNode) map[string]struct{} {
		set := make(map[string]struct{}, len(trees))
		for _, tree := range trees {
			serialized := tree.String()
			if _, exists := set[serialized]; exists {
				t.Errorf("duplicated %v", trees)
			}
			set[serialized] = struct{}{}
		}
		return set
	}
	for _, tt := range tests {
		assert.Equal(t, slice2Set(tt.want), slice2Set(GenerateTrees(tt.n)))
	}
}
