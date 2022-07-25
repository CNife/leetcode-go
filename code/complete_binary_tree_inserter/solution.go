package complete_binary_tree_inserter

import . "github.com/CNife/leetcode-go/types"

type CBTInserter struct {
	root  *TreeNode
	queue []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	if root == nil {
		panic("invalid root")
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left == nil || node.Right == nil {
			break
		} else {
			queue = queue[1:]
		}
	}
	return CBTInserter{root, queue}
}

func (inserter *CBTInserter) Insert(value int) int {
	if len(inserter.queue) < 1 {
		panic("invalid state")
	}

	parentNode := inserter.queue[0]
	newNode := &TreeNode{Val: value}
	if parentNode.Left == nil {
		parentNode.Left = newNode
	} else {
		parentNode.Right = newNode
		inserter.queue = inserter.queue[1:]
	}
	inserter.queue = append(inserter.queue, newNode)
	return parentNode.Val
}

func (inserter *CBTInserter) GetRoot() *TreeNode {
	return inserter.root
}
