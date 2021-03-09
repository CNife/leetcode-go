package sliding_window_median

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedianSlidingWindow(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []float64
	}{
		{
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			want: []float64{1, -1, -1, 3, 5, 6},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MedianSlidingWindow(tt.nums, tt.k))
	}
}
