package power_of_two

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfTwo(t *testing.T) {
	assert.Equal(t, true, IsPowerOfTwo(1))
	assert.Equal(t, true, IsPowerOfTwo(16))
	assert.Equal(t, false, IsPowerOfTwo(3))
	assert.Equal(t, true, IsPowerOfTwo(4))
	assert.Equal(t, false, IsPowerOfTwo(5))
}
