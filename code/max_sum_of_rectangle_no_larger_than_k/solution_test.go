package max_sum_of_rectangle_no_larger_than_k

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSumSubMatrix(t *testing.T) {
	assert.Equal(t, 2, MaxSumSubMatrix([][]int{{1, 0, 1}, {0, -2, 3}}, 2))
	assert.Equal(t, 3, MaxSumSubMatrix([][]int{{2, 2, -1}}, 3))
}
