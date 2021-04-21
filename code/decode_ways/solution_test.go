package decode_ways

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumDecodings(t *testing.T) {
	assert.Equal(t, 2, NumDecodings("12"))
	assert.Equal(t, 3, NumDecodings("226"))
	assert.Equal(t, 0, NumDecodings("0"))
	assert.Equal(t, 0, NumDecodings("06"))
}
