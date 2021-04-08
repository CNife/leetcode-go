package find_minimum_in_rotated_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMin(t *testing.T) {
	assert.Equal(t, 1, FindMin([]int{3, 4, 5, 1, 2}))
	assert.Equal(t, 0, FindMin([]int{4, 5, 6, 7, 0, 1, 2}))
	assert.Equal(t, 11, FindMin([]int{11, 13, 15, 17}))
}
