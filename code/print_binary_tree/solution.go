package print_binary_tree

import (
	"strconv"

	. "github.com/CNife/leetcode-go/types"
)

func PrintTree(root *TreeNode) [][]string {
	height := getTreeHeight(root)

	m, n := height, (1<<height)-1
	result := make([][]string, m)
	for i := range result {
		result[i] = make([]string, n)
		for j := range result[i] {
			result[i][j] = ""
		}
	}

	putValue(root, result, 0, (n-1)/2)
	return result
}

func getTreeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight, rightHeight := getTreeHeight(root.Left), getTreeHeight(root.Right)
	if leftHeight < rightHeight {
		return rightHeight + 1
	}
	return leftHeight + 1
}

func putValue(node *TreeNode, result [][]string, r, c int) {
	result[r][c] = strconv.Itoa(node.Val)
	if node.Left != nil {
		putValue(node.Left, result, r+1, c-1<<(len(result)-2-r))
	}
	if node.Right != nil {
		putValue(node.Right, result, r+1, c+1<<(len(result)-2-r))
	}
}
