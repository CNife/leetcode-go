package find_first_and_last_position_of_element_in_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		{
			nums:   nil,
			target: 0,
			want:   []int{-1, -1},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, SearchRange(tt.nums, tt.target))
	}
}
