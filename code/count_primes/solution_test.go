package count_primes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPrimes(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{
			n:    0,
			want: 0,
		},
		{
			n:    1,
			want: 0,
		},
		{
			n:    2,
			want: 0,
		},
		{
			n:    3,
			want: 1,
		},
		{
			n:    10,
			want: 4,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, CountPrimes(tt.n))
	}
}
