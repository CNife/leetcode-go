package sum_of_square_numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJudgeSquareSum(t *testing.T) {
	assert.Equal(t, true, JudgeSquareSum(1))
	assert.Equal(t, true, JudgeSquareSum(2))
	assert.Equal(t, false, JudgeSquareSum(3))
	assert.Equal(t, true, JudgeSquareSum(4))
	assert.Equal(t, true, JudgeSquareSum(5))
}
