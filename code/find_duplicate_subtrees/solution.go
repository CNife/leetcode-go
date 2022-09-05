package find_duplicate_subtrees

import . "github.com/CNife/leetcode-go/types"

func FindDuplicateSubtrees(root *TreeNode) []*TreeNode {
	subtrees := make(map[treeHash]subtree)
	repeatSubtrees := make(map[*TreeNode]struct{})
	idCounter := 0
	_ = dfs(root, subtrees, repeatSubtrees, &idCounter)

	result := make([]*TreeNode, 0, len(repeatSubtrees))
	for subtree := range repeatSubtrees {
		result = append(result, subtree)
	}
	return result
}

func dfs(root *TreeNode, subtrees map[treeHash]subtree, repeatSubTrees map[*TreeNode]struct{}, idCounter *int) int {
	if root == nil {
		return 0
	}
	hash := treeHash{
		rootVal:   root.Val,
		leftHash:  dfs(root.Left, subtrees, repeatSubTrees, idCounter),
		rightHash: dfs(root.Right, subtrees, repeatSubTrees, idCounter),
	}
	if tree, exists := subtrees[hash]; exists {
		repeatSubTrees[tree.root] = struct{}{}
		return tree.id
	} else {
		*idCounter++
		subtrees[hash] = subtree{
			id:   *idCounter,
			root: root,
		}
		return *idCounter
	}
}

type treeHash struct {
	rootVal, leftHash, rightHash int
}

type subtree struct {
	id   int
	root *TreeNode
}
