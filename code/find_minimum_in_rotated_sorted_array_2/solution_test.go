package find_minimum_in_rotated_sorted_array_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMin(t *testing.T) {
	assert.Equal(t, 1, FindMin([]int{1, 3, 5}))
	assert.Equal(t, 0, FindMin([]int{2, 2, 2, 0, 1}))
}
