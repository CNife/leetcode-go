package minimum_path_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			want: 7,
		},
		{
			grid: [][]int{
				{1, 2, 5},
				{3, 2, 1},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MinPathSum(tt.grid))
	}
}
