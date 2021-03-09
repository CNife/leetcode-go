package max_consecutive_ones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 1, 0, 1, 1, 1},
			want: 3,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, FindMaxConsecutiveOnes(tt.nums))
	}
}
