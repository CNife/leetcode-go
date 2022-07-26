package design_skiplist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipList(t *testing.T) {
	sl := Constructor()
	sl.Add(1)
	sl.Add(2)
	sl.Add(3)
	assert.False(t, sl.Search(0))
	sl.Add(4)
	assert.True(t, sl.Search(1))
	assert.False(t, sl.Erase(0))
	assert.True(t, sl.Erase(1))
	assert.False(t, sl.Search(1))
}
