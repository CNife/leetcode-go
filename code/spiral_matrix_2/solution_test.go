package spiral_matrix_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMatrix(t *testing.T) {
	tests := []struct {
		n    int
		want [][]int
	}{
		{
			n: 3,
			want: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			n: 1,
			want: [][]int{
				{1},
			},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, GenerateMatrix(tt.n))
	}
}
