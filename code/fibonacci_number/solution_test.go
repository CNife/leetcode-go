package fibonacci_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{
			n:    0,
			want: 0,
		},
		{
			n:    1,
			want: 1,
		},
		{
			n:    2,
			want: 1,
		},
		{
			n:    3,
			want: 2,
		},
		{
			n:    4,
			want: 3,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, Fib(tt.n))
	}
}
