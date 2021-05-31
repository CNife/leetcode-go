package power_of_four

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfFour(t *testing.T) {
	assert.Equal(t, true, IsPowerOfFour(16))
	assert.Equal(t, false, IsPowerOfFour(5))
	assert.Equal(t, true, IsPowerOfFour(1))
	assert.Equal(t, false, IsPowerOfFour(0))
}
