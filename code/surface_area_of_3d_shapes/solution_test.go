package surface_area_of_3d_shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSurfaceArea(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{
				{2},
			},
			want: 10,
		},
		{
			grid: [][]int{
				{1, 2},
				{3, 4},
			},
			want: 34,
		},
		{
			grid: [][]int{
				{1, 0},
				{0, 2},
			},
			want: 16,
		},
		{
			grid: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			want: 32,
		},
		{
			grid: [][]int{
				{2, 2, 2},
				{2, 1, 2},
				{2, 2, 2},
			},
			want: 46,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, SurfaceArea(tt.grid))
	}
}
