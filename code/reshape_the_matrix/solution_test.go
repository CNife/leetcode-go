package reshape_the_matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrixReshape(t *testing.T) {
	tests := []struct {
		nums [][]int
		r, c int
		want [][]int
	}{
		{
			nums: [][]int{
				{1, 2},
				{3, 4},
			},
			r: 1, c: 4,
			want: [][]int{
				{1, 2, 3, 4},
			},
		},
		{
			nums: [][]int{
				{1, 2},
				{3, 4},
			},
			r: 2, c: 4,
			want: [][]int{
				{1, 2},
				{3, 4},
			},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MatrixReshape(tt.nums, tt.r, tt.c))
	}
}
