package ugly_number_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthUglyNumber(t *testing.T) {
	uglyNumbers := []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12}
	for i := range uglyNumbers {
		assert.Equal(t, uglyNumbers[i], NthUglyNumber(i+1))
	}
}
