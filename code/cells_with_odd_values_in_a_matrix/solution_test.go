package cells_with_odd_values_in_a_matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOddCells(t *testing.T) {
	tests := []struct {
		m, n    int
		indices [][]int
		want    int
	}{
		{
			m: 2,
			n: 3,
			indices: [][]int{
				{0, 1},
				{1, 1},
			},
			want: 6,
		},
		{
			m: 2,
			n: 2,
			indices: [][]int{
				{1, 1},
				{0, 0},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, OddCells(tt.m, tt.n, tt.indices))
	}
}
