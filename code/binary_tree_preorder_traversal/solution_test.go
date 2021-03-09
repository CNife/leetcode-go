package binary_tree_preorder_traversal

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestPreorderTraversal(t *testing.T) {
	tests := []*types.TreeNode{
		nil,
		types.NewTree(1, -1, 2, 3),
	}
	for _, tt := range tests {
		assert.Equal(t, PreorderTraversalIteration(tt),
			PreorderTraversalRecursion(tt))
	}
}
