package bricks_falling_when_hit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHitBricks(t *testing.T) {
	tests := []struct {
		grid [][]int
		hits [][]int
		want []int
	}{
		{
			grid: [][]int{
				{1, 0, 0, 0},
				{1, 1, 1, 0},
			},
			hits: [][]int{
				{1, 0},
			},
			want: []int{2},
		},
		{
			grid: [][]int{
				{1, 0, 0, 0},
				{1, 1, 0, 0},
			},
			hits: [][]int{
				{1, 1},
				{1, 0},
			},
			want: []int{0, 0},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, HitBricks(tt.grid, tt.hits))
	}
}
