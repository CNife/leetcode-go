package next_permutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		nums, want []int
	}{
		{
			nums: []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			nums: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		{
			nums: []int{1, 1, 5},
			want: []int{1, 5, 1},
		},
	}
	for _, tt := range tests {
		NextPermutation(tt.nums)
		assert.Equal(t, tt.want, tt.nums)
	}
}
