package rotting_oranges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			want: 4,
		},
		{
			grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			want: -1,
		},
		{
			grid: [][]int{
				{0, 2},
			},
			want: 0,
		},
		{
			grid: [][]int{
				{0},
			},
			want: 0,
		},
		{
			grid: [][]int{
				{1},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, OrangesRotting(tt.grid))
	}
}
