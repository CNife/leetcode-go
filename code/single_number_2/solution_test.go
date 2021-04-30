package single_number_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	assert.Equal(t, 3, SingleNumber([]int{2, 2, 3, 2}))
	assert.Equal(t, 99, SingleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
}
