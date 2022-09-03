package longest_univalue_path

import . "github.com/CNife/leetcode-go/types"

func LongestUnivaluePath(root *TreeNode) int {
	result := 0
	_ = arrowLength(root, &result)
	return result
}

func arrowLength(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	left, leftPath := arrowLength(root.Left, result), 0
	if root.Left != nil && root.Val == root.Left.Val {
		leftPath += left + 1
	}
	right, rightPath := arrowLength(root.Right, result), 0
	if root.Right != nil && root.Val == root.Right.Val {
		rightPath += right + 1
	}
	if pathSum := leftPath + rightPath; *result < pathSum {
		*result = pathSum
	}
	if leftPath < rightPath {
		return rightPath
	}
	return leftPath
}
