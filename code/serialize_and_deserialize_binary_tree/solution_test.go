package serialize_and_deserialize_binary_tree

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestCodec(t *testing.T) {
	trees := []*types.TreeNode{
		types.NewTree(),
		types.NewTree(1, 2, 3, 4, 5, 6, 7),
		types.NewTree(1, -1, 2, -1, 3, -1, 4),
	}
	c := Constructor()
	for _, tree := range trees {
		assert.Equal(t, tree, c.Deserialize(c.Serialize(tree)))
	}
}
