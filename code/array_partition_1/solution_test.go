package array_partition_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayPairSum(t *testing.T) {
	assert.Equal(t, 4, ArrayPairSum([]int{1, 4, 3, 2}))
	assert.Equal(t, 9, ArrayPairSum([]int{6, 2, 6, 5, 1, 2}))
}
