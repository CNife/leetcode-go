package frog_jump

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanCross(t *testing.T) {
	assert.Equal(t, true, CanCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
	assert.Equal(t, false, CanCross([]int{0, 1, 2, 3, 4, 8, 9, 11}))
}
