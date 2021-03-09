package number_of_equivalent_domino_pairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumEquivDominoPairs(t *testing.T) {
	tests := []struct {
		dominoes [][]int
		want     int
	}{
		{
			dominoes: [][]int{
				{1, 2},
				{2, 1},
				{3, 4},
				{5, 6},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, NumEquivDominoPairs(tt.dominoes))
	}
}
