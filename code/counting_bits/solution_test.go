package counting_bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBits(t *testing.T) {
	bits := []int{0, 1, 1, 2, 1, 2}
	for i := 0; i < len(bits); i++ {
		assert.Equal(t, bits[:i+1], CountBits(i))
	}
}
