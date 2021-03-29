package reverse_bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBits(t *testing.T) {
	tests := []struct {
		num, want uint32
	}{
		{
			num:  0b00000010100101000001111010011100,
			want: 0b00111001011110000010100101000000,
		},
		{
			num:  0b11111111111111111111111111111101,
			want: 0b10111111111111111111111111111111,
		},
		{
			num:  0b00000010100101000001111010011100,
			want: 0b00111001011110000010100101000000,
		},
		{
			num:  0b11111111111111111111111111111101,
			want: 0b10111111111111111111111111111111,
		},
		{
			num:  0b00000010100101000001111010011100,
			want: 0b00111001011110000010100101000000,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, ReverseBits(tt.num))
	}
}
