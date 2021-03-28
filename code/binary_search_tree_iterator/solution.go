package binary_search_tree_iterator

import "github.com/CNife/leetcode-go/types"

type BSTIterator struct {
	stack []stackItem
}

func Constructor(root *types.TreeNode) BSTIterator {
	return BSTIterator{
		stack: []stackItem{{root, white}},
	}
}

func (bi *BSTIterator) HasNext() bool {
	return len(bi.stack) > 0
}

func (bi *BSTIterator) Next() int {
	for {
		item := bi.stack[len(bi.stack)-1]
		bi.stack = bi.stack[:len(bi.stack)-1]
		node, color := item.node, item.color

		if color == gray {
			return node.Val
		}
		if node.Right != nil {
			bi.stack = append(bi.stack, stackItem{node.Right, white})
		}
		bi.stack = append(bi.stack, stackItem{node, gray})
		if node.Left != nil {
			bi.stack = append(bi.stack, stackItem{node.Left, white})
		}
	}
}

type stackItem struct {
	node  *types.TreeNode
	color byte
}

const (
	white byte = iota + 1
	gray
)
