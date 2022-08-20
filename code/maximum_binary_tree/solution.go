package maximum_binary_tree

import . "github.com/CNife/leetcode-go/types"

func ConstructMaximumBinaryTree(nums []int) *TreeNode {
	var stack []*TreeNode
	for _, num := range nums {
		node := &TreeNode{
			Val: num,
		}
		for len(stack) > 0 {
			topNode := stack[len(stack)-1]
			if topNode.Val > num {
				stack = append(stack, node)
				topNode.Right = node
				break
			} else {
				stack = stack[:len(stack)-1]
				node.Left = topNode
			}
		}
		if len(stack) < 1 {
			stack = append(stack, node)
		}
	}
	return stack[0]
}
