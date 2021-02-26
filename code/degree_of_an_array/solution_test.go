package degree_of_an_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindShortestSubArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 2, 2, 3, 1},
			want: 2,
		},
		{
			nums: []int{1, 2, 2, 3, 1, 4, 2},
			want: 6,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, FindShortestSubArray(tt.nums))
	}
}
