package house_robber_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRob(t *testing.T) {
	assert.Equal(t, 3, Rob([]int{2, 3, 2}))
	assert.Equal(t, 4, Rob([]int{1, 2, 3, 1}))
	assert.Equal(t, 0, Rob([]int{0}))
}
