package number_of_1_bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		num  uint32
		want int
	}{
		{
			num:  0b00000000000000000000000000001011,
			want: 3,
		},
		{
			num:  0b00000000000000000000000010000000,
			want: 1,
		},
		{
			num:  0b11111111111111111111111111111101,
			want: 31,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, HammingWeight(tt.num))
	}
}
