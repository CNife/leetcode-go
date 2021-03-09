package sum_root_to_leaf_numbers

import "github.com/CNife/leetcode-go/types"

func SumNumbers(root *types.TreeNode) int {
	if root == nil {
		return 0
	}

	num, result := 0, 0

	var dfs func(node *types.TreeNode)
	dfs = func(node *types.TreeNode) {
		num = num*10 + node.Val
		if node.Left == nil && node.Right == nil {
			result += num
		} else {
			if node.Left != nil {
				dfs(node.Left)
			}
			if node.Right != nil {
				dfs(node.Right)
			}
		}
		num /= 10
	}

	dfs(root)
	return result
}
