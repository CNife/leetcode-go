package my_calendar_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyCalendarTwo(t *testing.T) {
	c := Constructor()
	assert.True(t, c.Book(10, 20))
	assert.True(t, c.Book(50, 60))
	assert.True(t, c.Book(10, 40))
	assert.False(t, c.Book(5, 15))
	assert.True(t, c.Book(5, 10))
	assert.True(t, c.Book(25, 55))
}
