package maximum_binary_tree_ii

import . "github.com/CNife/leetcode-go/types"

func InsertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{
		Val: val,
	}
	if root == nil {
		return newNode
	}
	if val > root.Val {
		newNode.Left = root
		return newNode
	}
	root.Right = InsertIntoMaxTree(root.Right, val)
	return root
}
