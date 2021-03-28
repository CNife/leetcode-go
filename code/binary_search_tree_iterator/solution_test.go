package binary_search_tree_iterator

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestBSTIterator(t *testing.T) {
	values := []int{7, 3, 15, -1, -1, 9, 20}
	sortedValues := []int{3, 7, 9, 15, 20}
	tree := types.NewTree(values...)

	var gotValues []int
	bi := Constructor(tree)
	for bi.HasNext() {
		gotValues = append(gotValues, bi.Next())
	}
	assert.Equal(t, sortedValues, gotValues)
}
