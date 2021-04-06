package remove_duplicates_from_sorted_array_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums, want []int
	}{
		{
			nums: []int{1, 1, 1, 2, 2, 3},
			want: []int{1, 1, 2, 2, 3},
		},
		{
			nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			want: []int{0, 0, 1, 1, 2, 3, 3},
		},
	}
	for _, tt := range tests {
		got := RemoveDuplicates(tt.nums)
		assert.Equal(t, tt.want, tt.nums[:got])
	}
}
