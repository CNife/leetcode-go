package split_array_into_consecutive_subsequences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPossible(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{1, 2, 3, 3, 4, 5},
			want: true,
		},
		{
			nums: []int{1, 2, 3, 3, 4, 4, 5, 5},
			want: true,
		},
		{
			nums: []int{1, 2, 3, 4, 4, 5},
			want: false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, IsPossible(tt.nums))
	}
}
