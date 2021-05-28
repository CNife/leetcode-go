package hamming_distance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHammingDistance(t *testing.T) {
	assert.Equal(t, 2, HammingDistance(1, 4))
}
