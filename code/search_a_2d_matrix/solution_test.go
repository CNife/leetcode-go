package search_a_2d_matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 3,
			want:   true,
		},
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 13,
			want:   false,
		},
		{
			matrix: [][]int{
				{1},
			},
			target: 1,
			want:   true,
		},
		{
			matrix: [][]int{
				{1},
			},
			target: 2,
			want:   false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, SearchMatrix(tt.matrix, tt.target))
	}
}
