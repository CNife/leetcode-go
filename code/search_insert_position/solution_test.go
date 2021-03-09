package search_insert_position

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums         []int
		target, want int
	}{
		{
			nums:   nil,
			target: 0,
			want:   0,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 0,
			want:   0,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, SearchInsert(tt.nums, tt.target))
	}
}
