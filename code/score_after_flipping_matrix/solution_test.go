package score_after_flipping_matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrixScore(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   int
	}{
		{
			matrix: [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}},
			want:   39,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MatrixScore(tt.matrix))
	}
}
