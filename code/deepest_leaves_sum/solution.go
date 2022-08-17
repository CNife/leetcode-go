package deepest_leaves_sum

import . "github.com/CNife/leetcode-go/types"

func DeepestLeaveSum(root *TreeNode) int {
	var prevLevel []*TreeNode
	currLevel := []*TreeNode{root}
	for len(currLevel) > 0 {
		prevLevel, currLevel = currLevel, make([]*TreeNode, 0, len(currLevel)*2)
		for _, node := range prevLevel {
			if node.Left != nil {
				currLevel = append(currLevel, node.Left)
			}
			if node.Right != nil {
				currLevel = append(currLevel, node.Right)
			}
		}
	}

	var result int
	for _, node := range prevLevel {
		result += node.Val
	}
	return result
}
