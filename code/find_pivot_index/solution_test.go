package find_pivot_index

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 7, 3, 6, 5, 6},
			want: 3,
		},
		{
			nums: []int{1, 2, 3},
			want: -1,
		},
		{
			nums: []int{-1, -1, -1, 0, 1, 1},
			want: 0,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, PivotIndex(tt.nums))
	}
}
