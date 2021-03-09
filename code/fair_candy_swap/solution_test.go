package fair_candy_swap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFairCandySwap(t *testing.T) {
	tests := []struct {
		a, b, want []int
	}{
		{
			a:    []int{1, 1},
			b:    []int{2, 2},
			want: []int{1, 2},
		},
		{
			a:    []int{1, 2},
			b:    []int{2, 3},
			want: []int{1, 2},
		},
		{
			a:    []int{2},
			b:    []int{1, 3},
			want: []int{2, 3},
		},
		{
			a:    []int{1, 2, 5},
			b:    []int{2, 4},
			want: []int{5, 4},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, FairCandySwap(tt.a, tt.b))
	}
}
