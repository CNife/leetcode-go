package summary_ranges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		nums []int
		want []string
	}{
		{
			nums: []int{0, 1, 2, 4, 5, 7},
			want: []string{"0->2", "4->5", "7"},
		},
		{
			nums: []int{0, 2, 3, 4, 6, 8, 9},
			want: []string{"0", "2->4", "6", "8->9"},
		},
		{
			nums: nil,
			want: nil,
		},
		{
			nums: []int{-1},
			want: []string{"-1"},
		},
		{
			nums: []int{0},
			want: []string{"0"},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, SummaryRanges(tt.nums))
	}
}
