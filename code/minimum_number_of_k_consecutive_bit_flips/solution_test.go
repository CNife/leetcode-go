package minimum_number_of_k_consecutive_bit_flips

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinKBitFlips(t *testing.T) {
	tests := []struct {
		array   []int
		k, want int
	}{
		{
			array: []int{0, 1, 0},
			k:     1,
			want:  2,
		},
		{
			array: []int{1, 1, 0},
			k:     2,
			want:  -1,
		},
		{
			array: []int{0, 0, 0, 1, 0, 1, 1, 0},
			k:     3,
			want:  3,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MinKBitFlips(tt.array, tt.k))
	}
}
