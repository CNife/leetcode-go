package range_sum_of_bst

import (
	"testing"

	"github.com/CNife/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestRangeSumBST(t *testing.T) {
	assert.Equal(t, 32,
		RangeSumBST(
			types.NewTree(10, 5, 15, 3, 7, -1, 18),
			7,
			15,
		),
	)
	assert.Equal(t, 23,
		RangeSumBST(
			types.NewTree(10, 5, 15, 3, 7, 13, 18, 1, -1, 6),
			6,
			10),
	)
}
