package reverse_pairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReversePairs(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 3, 2, 3, 1},
			want: 2,
		},
		{
			nums: []int{2, 4, 3, 5, 1},
			want: 3,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, ReversePairs(tt.nums))
	}
}
