package how_many_numbers_are_smaller_than_the_current_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallerNumbersThanCurrent(t *testing.T) {
	tests := []struct {
		nums, want []int
	}{
		{
			nums: []int{8, 1, 2, 2, 3},
			want: []int{4, 0, 1, 1, 3},
		},
		{
			nums: []int{6, 5, 4, 8},
			want: []int{2, 1, 0, 3},
		},
		{
			nums: []int{7, 7, 7, 7},
			want: []int{0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, SmallerNumbersThanCurrent(tt.nums))
	}
}
