package first_missing_positive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 2, 0},
			want: 3,
		},
		{
			nums: []int{3, 4, 1, -1},
			want: 2,
		},
		{
			nums: []int{7, 8, 9, 11, 12},
			want: 1,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, FirstMissingPositive(tt.nums))
	}
}
