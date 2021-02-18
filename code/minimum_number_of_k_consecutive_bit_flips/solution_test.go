package minimum_number_of_k_consecutive_bit_flips

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinKBitFlips(t *testing.T) {
	assert.Equal(t, 2, MinKBitFlips([]int{0, 1, 0}, 1))
	assert.Equal(t, -1, MinKBitFlips([]int{1, 1, 0}, 2))
	assert.Equal(t, 3, MinKBitFlips([]int{0, 0, 0, 1, 0, 1, 1, 0}, 3))
}
