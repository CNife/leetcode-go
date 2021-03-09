package subarrays_with_k_different_integers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubarraysWithKDistinct(t *testing.T) {
	tests := []struct {
		nums    []int
		k, want int
	}{
		{
			nums: []int{1, 2, 1, 2, 3}, k: 2,
			want: 7,
		},
		{
			nums: []int{1, 2, 1, 3, 4}, k: 3,
			want: 3,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, SubarraysWithKDistinct(tt.nums, tt.k))
	}
}
