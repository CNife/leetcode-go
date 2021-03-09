package wiggle_subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWiggleMaxLength(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 7, 4, 9, 2, 5},
			want: 6,
		},
		{
			nums: []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8},
			want: 7,
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 2,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, WiggleMaxLength(tt.nums))
	}
}
