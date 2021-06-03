package contiguous_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxLength(t *testing.T) {
	assert.Equal(t, 2, FindMaxLength([]int{0, 1}))
	assert.Equal(t, 2, FindMaxLength([]int{0, 1, 0}))
}
