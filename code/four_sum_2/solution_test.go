package four_sum_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFourSumCount(t *testing.T) {
	tests := []struct {
		a, b, c, d []int
		want       int
	}{
		{
			a:    []int{1, 2},
			b:    []int{-2, -1},
			c:    []int{-1, 2},
			d:    []int{0, 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, FourSumCount(tt.a, tt.b, tt.c, tt.d))
	}
}
