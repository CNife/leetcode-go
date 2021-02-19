package max_consecutive_ones_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestOnes(t *testing.T) {
	tests := []struct {
		nums    []int
		k, want int
	}{
		{
			nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			k:    2,
			want: 6,
		},
		{
			nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			k:    3,
			want: 10,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, LongestOnes(tt.nums, tt.k))
	}
}
