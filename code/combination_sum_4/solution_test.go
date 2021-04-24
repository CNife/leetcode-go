package combination_sum_4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum4(t *testing.T) {
	assert.Equal(t, 7, CombinationSum4([]int{1, 2, 3}, 4))
	assert.Equal(t, 0, CombinationSum4([]int{9}, 3))
	assert.Equal(t, 181997601, CombinationSum4([]int{1, 2, 3}, 32))
}
