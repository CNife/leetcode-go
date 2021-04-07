package search_in_rotated_sorted_array_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   bool
	}{
		{
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 0,
			want:   true,
		},
		{
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 3,
			want:   false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, Search(tt.nums, tt.target))
	}
}
