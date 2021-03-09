package array_partition_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayPairSum(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 4, 3, 2},
			want: 4,
		},
		{
			nums: []int{6, 2, 6, 5, 1, 2},
			want: 9,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, ArrayPairSum(tt.nums))
	}
}
