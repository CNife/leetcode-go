package largest_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestNumber(t *testing.T) {
	assert.Equal(t, "210", LargestNumber([]int{10, 2}))
	assert.Equal(t, "9534330", LargestNumber([]int{3, 30, 34, 5, 9}))
	assert.Equal(t, "1", LargestNumber([]int{1}))
	assert.Equal(t, "10", LargestNumber([]int{10}))
	assert.Equal(t, "0", LargestNumber([]int{0, 0}))
}
