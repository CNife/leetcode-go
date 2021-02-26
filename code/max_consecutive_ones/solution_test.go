package max_consecutive_ones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	assert.Equal(t, 3, FindMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}
