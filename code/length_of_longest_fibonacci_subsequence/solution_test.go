package length_of_longest_fibonacci_subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLenLongestFibSubSeq(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8},
			want: 5,
		},
		{
			nums: []int{1, 3, 7, 11, 12, 14, 18},
			want: 3,
		},
		{
			nums: []int{1, 3, 5},
			want: 0,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, LenLongestFibSubSeq(tt.nums))
	}
}
