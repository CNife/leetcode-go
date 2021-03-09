package count_of_range_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountRangeSum(t *testing.T) {
	tests := []struct {
		nums          []int
		lower, higher int
		want          int
	}{
		{
			nums:   []int{-2, 5, -1},
			lower:  -2,
			higher: 2,
			want:   3,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, CountRangeSum(tt.nums, tt.lower, tt.higher))
	}
}
