package increasing_order_search_tree

import "github.com/CNife/leetcode-go/types"

func IncreasingBST(root *types.TreeNode) *types.TreeNode {
	var stack []*types.TreeNode
	fakeRoot := types.TreeNode{Val: -1}
	node := &fakeRoot
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if n := len(stack); n > 0 {
			root = stack[n-1]
			stack = stack[:n-1]
			node.Right = root
			node = root
			root.Left = nil
			root = root.Right
		}
	}
	node.Right = nil
	return fakeRoot.Right
}
