package coin_change_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChange(t *testing.T) {
	assert.Equal(t, 4, Change(5, []int{1, 2, 5}))
	assert.Equal(t, 0, Change(3, []int{2}))
	assert.Equal(t, 1, Change(10, []int{10}))
}
