package my_calendar_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyCalendar(t *testing.T) {
	c := Constructor()
	assert.True(t, c.Book(10, 20))
	assert.False(t, c.Book(15, 25))
	assert.True(t, c.Book(20, 30))

	c.Clear()
	assert.True(t, c.Book(37, 50))
	assert.False(t, c.Book(33, 50))
	assert.True(t, c.Book(4, 17))
	assert.False(t, c.Book(35, 48))
	assert.False(t, c.Book(8, 25))
}
