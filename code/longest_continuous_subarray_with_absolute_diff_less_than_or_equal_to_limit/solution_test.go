package longest_continuous_subarray_with_absolute_diff_less_than_or_equal_to_limit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		nums        []int
		limit, want int
	}{
		{
			nums:  []int{8, 2, 4, 7},
			limit: 4,
			want:  2,
		},
		{
			nums:  []int{10, 1, 2, 4, 7, 2},
			limit: 5,
			want:  4,
		},
		{
			nums:  []int{4, 2, 2, 2, 4, 4, 2, 2},
			limit: 0,
			want:  3,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, LongestSubarray(tt.nums, tt.limit))
	}
}
