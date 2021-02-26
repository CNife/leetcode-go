package array_partition_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayPairSum(t *testing.T) {
	assert.Equal(t, 4, ArrayPairSum([]int{1, 4, 3, 2}))
	assert.Equal(t, 9, ArrayPairSum([]int{6, 2, 6, 5, 1, 2}))
}
