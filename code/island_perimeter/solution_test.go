package island_perimeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIslandPerimeter(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{
				{0, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 1, 0, 0},
				{1, 1, 0, 0},
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, IslandPerimeter(tt.grid))
	}
}
