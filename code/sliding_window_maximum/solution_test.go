package sliding_window_maximum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			nums: []int{1, -1},
			k:    1,
			want: []int{1, -1},
		},
		{
			nums: []int{9, 11},
			k:    2,
			want: []int{11},
		},
		{
			nums: []int{4, -2},
			k:    2,
			want: []int{4},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MaxSlidingWindow(tt.nums, tt.k))
	}
}
