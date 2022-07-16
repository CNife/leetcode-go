package average_of_sliding_window

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingAverage(t *testing.T) {
	mv := Constructor(3)
	assert.InEpsilon(t, 1.0, mv.Next(1), 1e-9)
	assert.InEpsilon(t, 5.5, mv.Next(10), 1e-9)
	assert.InEpsilon(t, 14.0/3, mv.Next(3), 1e-9)
	assert.InEpsilon(t, 6.0, mv.Next(5), 1e-9)
}
