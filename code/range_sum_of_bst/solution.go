package range_sum_of_bst

import . "github.com/CNife/leetcode-go/types"

func RangeSumBST(root *TreeNode, low, high int) int {
	sum := 0

	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}

		val := node.Val
		if val < low {
			helper(node.Right)
		} else if val > high {
			helper(node.Left)
		} else {
			sum += val
			helper(node.Left)
			helper(node.Right)
		}
	}
	helper(root)

	return sum
}
