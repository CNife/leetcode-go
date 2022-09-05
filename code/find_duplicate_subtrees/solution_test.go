package find_duplicate_subtrees

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
)

func TestFindDuplicateSubtrees(t *testing.T) {
	tests := []struct {
		root *types.TreeNode
		want []*types.TreeNode
	}{
		{
			root: types.NewTree(1, 2, 3, 4, -1, 2, 4, -1, -1, 4),
			want: []*types.TreeNode{
				types.NewTree(2, 4),
				types.NewTree(4),
			},
		},
		{
			root: types.NewTree(2, 1, 1),
			want: []*types.TreeNode{
				types.NewTree(1),
			},
		},
		{
			root: types.NewTree(2, 2, 2, 3, -1, 3, -1),
			want: []*types.TreeNode{
				types.NewTree(2, 3),
				types.NewTree(3),
			},
		},
	}
	for _, tt := range tests {
		// assert.Equal(t, makeTreesMap(tt.want), makeTreesMap(FindDuplicateSubtrees(tt.root)))
		_ = FindDuplicateSubtrees(tt.root)
	}
}

// func makeTreesMap(trees []*types.TreeNode) map[*types.TreeNode]struct{} {
// 	treesMap := make(map[*types.TreeNode]struct{}, len(trees))
// 	for _, tree := range trees {
// 		treesMap[tree] = struct{}{}
// 	}
// 	return treesMap
// }
