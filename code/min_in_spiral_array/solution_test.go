package min_in_spiral_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			nums: []int{2, 2, 2, 0, 1},
			want: 0,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MinArray(tt.nums))
	}
}
