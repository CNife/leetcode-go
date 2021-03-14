package design_hashmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyHashMap(t *testing.T) {
	h := Constructor()
	h.Put(1, 1)
	h.Put(2, 2)
	assert.Equal(t, 1, h.Get(1))
	assert.Equal(t, -1, h.Get(3))
	h.Put(2, 1)
	assert.Equal(t, 1, h.Get(2))
	h.Remove(2)
	assert.Equal(t, -1, h.Get(2))
}
