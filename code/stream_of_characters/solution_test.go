package stream_of_characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStreamChecker(t *testing.T) {
	sc := Constructor([]string{"cd", "f", "kl"})
	assert.False(t, sc.Query('a'))
	assert.False(t, sc.Query('b'))
	assert.False(t, sc.Query('c'))
	assert.True(t, sc.Query('d'))
	assert.False(t, sc.Query('e'))
	assert.True(t, sc.Query('f'))
	assert.False(t, sc.Query('g'))
	assert.False(t, sc.Query('h'))
	assert.False(t, sc.Query('i'))
	assert.False(t, sc.Query('j'))
	assert.False(t, sc.Query('k'))
	assert.True(t, sc.Query('l'))
}
