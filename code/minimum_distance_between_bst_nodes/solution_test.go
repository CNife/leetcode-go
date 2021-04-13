package minimum_distance_between_bst_nodes

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestMinDiffInBST(t *testing.T) {
	assert.Equal(t, 1, MinDiffInBST(types.NewTree(4, 2, 6, 1, 3)))
	assert.Equal(t, 1, MinDiffInBST(types.NewTree(1, 0, 48, -1, -1, 12, 49)))
}
