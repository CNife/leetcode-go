package maximum_gap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumGap(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: nil,
			want: 0,
		},
		{
			nums: []int{1},
			want: 0,
		},
		{
			nums: []int{1, 10000000},
			want: 9999999,
		},
		{
			nums: []int{3, 6, 9, 1},
			want: 3,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MaximumGap(tt.nums))
	}
}
