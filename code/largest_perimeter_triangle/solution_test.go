package largest_perimeter_triangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestPerimeter(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{2, 1, 2},
			want: 5,
		},
		{
			nums: []int{1, 2, 1},
			want: 0,
		},
		{
			nums: []int{3, 2, 3, 4},
			want: 10,
		},
		{
			nums: []int{3, 6, 2, 3},
			want: 8,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, LargestPerimeter(tt.nums))
	}
}
