package split_array_into_fibonacci_sequence

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitIntoFibonacci(t *testing.T) {
	tests := []struct {
		s           string
		hasSolution bool
	}{
		{
			s:           "0123",
			hasSolution: false,
		},
		{
			s:           "123456579",
			hasSolution: true,
		},
		{
			s:           "11235813",
			hasSolution: true,
		},
		{
			s:           "112358130",
			hasSolution: false,
		},
		{
			s:           "1101111",
			hasSolution: true,
		},
	}
	for _, tt := range tests {
		assert.True(t,
			isValidSolution(tt.s, tt.hasSolution, SplitIntoFibonacci(tt.s)))
	}
}

func isValidSolution(s string, hasSolution bool, got []int) bool {
	if !hasSolution {
		return len(got) == 0
	} else if len(got) == 0 {
		return false
	}
	for i := 0; i < len(got)-2; i++ {
		if got[i]+got[i+1] != got[i+2] {
			return false
		}
	}
	i := 0
	for _, num := range got {
		numString := strconv.Itoa(num)
		endI := i + len(numString)
		if endI > len(s) || s[i:endI] != numString {
			return false
		}
		i = endI
	}
	return true
}
