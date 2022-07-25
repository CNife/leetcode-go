package complete_binary_tree_inserter

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestCBTInserter(t *testing.T) {
	inserter := Constructor(types.NewTree(1, 2))
	assert.Equal(t, 1, inserter.Insert(3))
	assert.Equal(t, 2, inserter.Insert(4))
	assert.Equal(t, types.NewTree(1, 2, 3, 4), inserter.GetRoot())
}
