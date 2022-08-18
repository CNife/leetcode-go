package maximum_equal_frequency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxEqualFreq(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{2, 2, 1, 1, 5, 3, 3, 5},
			want: 7,
		},
		{
			nums: []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5},
			want: 13,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MaxEqualFreq(tt.nums))
	}
}
