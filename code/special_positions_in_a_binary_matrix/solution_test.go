package special_positions_in_a_binary_matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSpecial(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   int
	}{
		{
			matrix: [][]int{
				{1, 0, 0},
				{0, 0, 1},
				{1, 0, 0},
			},
			want: 1,
		},
		{
			matrix: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			want: 3,
		},
		{
			matrix: [][]int{
				{0, 0, 0, 1},
				{1, 0, 0, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 0},
			},
			want: 2,
		},
		{
			matrix: [][]int{
				{0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{0, 1, 0, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 0, 1, 1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, NumSpecial(tt.matrix))
	}
}
