package maximum_length_of_pair_chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLongestChain(t *testing.T) {
	tests := []struct {
		pairs [][]int
		want  int
	}{
		{
			pairs: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
			},
			want: 2,
		},
		{
			pairs: [][]int{
				{-10, -8},
				{8, 9},
				{-5, 0},
				{6, 10},
				{-6, -4},
				{1, 7},
				{9, 10},
				{-4, 7},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, FindLongestChain(tt.pairs))
	}
}
