package design_an_ordered_stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderedStream(t *testing.T) {
	s := Constructor(5)
	assert.Equal(t, 0, len(s.Insert(3, "ccccc")))
	assert.Equal(t, []string{"aaaaa"}, s.Insert(1, "aaaaa"))
	assert.Equal(t, []string{"bbbbb", "ccccc"}, s.Insert(2, "bbbbb"))
	assert.Equal(t, 0, len(s.Insert(5, "eeeee")))
	assert.Equal(t, []string{"ddddd", "eeeee"}, s.Insert(4, "ddddd"))
}
