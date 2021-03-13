package design_hashset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyHashSet(t *testing.T) {
	h := Constructor()
	h.Add(1)
	h.Add(2)
	assert.True(t, h.Contains(1))
	assert.False(t, h.Contains(3))
	h.Add(2)
	assert.True(t, h.Contains(2))
	h.Remove(2)
	assert.False(t, h.Contains(2))
}
