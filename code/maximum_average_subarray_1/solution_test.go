package maximum_average_subarray_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want float64
	}{
		{
			nums: []int{1, 12, -5, -6, 50, 3},
			k:    4,
			want: 12.75,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, FindMaxAverage(tt.nums, tt.k))
	}
}
