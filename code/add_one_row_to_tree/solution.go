package add_one_row_to_tree

import . "github.com/CNife/leetcode-go/types"

func AddOneRow(root *TreeNode, val, depth int) *TreeNode {
	if depth < 1 {
		panic("invalid depth")
	}
	return addNode(root, val, depth, true)
}

func addNode(node *TreeNode, val, depth int, isLeft bool) *TreeNode {
	if depth == 1 {
		newRoot := &TreeNode{Val: val}
		if isLeft {
			newRoot.Left = node
		} else {
			newRoot.Right = node
		}
		return newRoot
	}
	if node != nil {
		node.Left = addNode(node.Left, val, depth-1, true)
		node.Right = addNode(node.Right, val, depth-1, false)
	}
	return node
}
