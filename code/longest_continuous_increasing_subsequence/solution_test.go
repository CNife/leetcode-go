package longest_continuous_increasing_subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLengthOfLCIS(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 3, 5, 4, 7},
			want: 3,
		},
		{
			nums: []int{2, 2, 2, 2, 2},
			want: 1,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, FindLengthOfLCIS(tt.nums))
	}
}
