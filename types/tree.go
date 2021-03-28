package types

import "strconv"

// 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 创建新的二叉树
// 按照值的层序遍历创建树，小于 0 的值视为空节点，空节点将在下一层中被忽略
func NewTree(values ...int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{Val: values[0]}
	queue := []*TreeNode{root}

	for i := 1; i < len(values); {
		node := queue[0]
		queue = queue[1:]

		leftValue := values[i]
		i++
		if leftValue >= 0 {
			node.Left = &TreeNode{Val: leftValue}
			queue = append(queue, node.Left)
		}

		if i < len(values) {
			rightValue := values[i]
			i++
			if rightValue >= 0 {
				node.Right = &TreeNode{Val: rightValue}
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}

// 打印二叉树
func (tree *TreeNode) String() string {
	if tree == nil {
		return "nil"
	}
	return "(" + strconv.Itoa(tree.Val) + "," +
		tree.Left.String() + "," +
		tree.Right.String() + ")"
}

// 克隆二叉树
func (tree *TreeNode) Clone() *TreeNode {
	if tree == nil {
		return nil
	}
	return &TreeNode{
		Val:   tree.Val,
		Left:  tree.Left.Clone(),
		Right: tree.Right.Clone(),
	}
}

func (tree *TreeNode) Len() int {
	if tree == nil {
		return 0
	}
	return 1 + tree.Left.Len() + tree.Right.Len()
}
