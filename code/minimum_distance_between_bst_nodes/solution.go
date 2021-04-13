package minimum_distance_between_bst_nodes

import (
	"math"

	"github.com/CNife/leetcode-go/types"
)

func MinDiffInBST(root *types.TreeNode) int {
	var stack []*types.TreeNode
	node, prevValue, result := root, -1, math.MaxInt32

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		diff := node.Val - prevValue
		if prevValue >= 0 && result > diff {
			result = diff
		}
		prevValue = node.Val

		node = node.Right
	}

	return result
}
