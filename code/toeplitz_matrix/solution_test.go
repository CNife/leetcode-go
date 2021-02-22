package toeplitz_matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsToeplitzMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   bool
	}{
		{
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 1, 2, 3},
				{9, 5, 1, 2},
			},
			want: true,
		},
		{
			matrix: [][]int{
				{1, 2},
				{2, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, IsToeplitzMatrix(tt.matrix))
	}
}
