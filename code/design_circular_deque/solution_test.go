package design_circular_deque

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyCircularDeque(t *testing.T) {
	deque := Constructor(3)
	assert.True(t, deque.InsertLast(1))
	assert.True(t, deque.InsertLast(2))
	assert.True(t, deque.InsertFront(3))
	assert.False(t, deque.InsertFront(3))
	assert.Equal(t, 2, deque.GetRear())
	assert.True(t, deque.IsFull())
	assert.True(t, deque.DeleteLast())
	assert.True(t, deque.InsertFront(4))
	assert.Equal(t, 4, deque.GetFront())
	assert.False(t, deque.IsEmpty())
}
