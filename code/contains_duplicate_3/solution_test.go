package contains_duplicate_3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		k, t int
		want bool
	}{
		{
			nums: []int{2147483647, -1, 2147483647},
			k:    1,
			t:    2147483647,
			want: false,
		},
		{
			nums: []int{1, 3, 6, 2},
			k:    1,
			t:    2,
			want: true,
		},
		{
			nums: []int{0, 10, 22, 15, 0, 5, 22, 12, 1, 5},
			k:    3,
			t:    3,
			want: false,
		},
		{
			nums: []int{1, 2, 3, 1},
			k:    3,
			t:    0,
			want: true,
		},
		{
			nums: []int{1, 0, 1, 1},
			k:    1,
			t:    2,
			want: true,
		},
		{
			nums: []int{1, 5, 9, 1, 5, 9},
			k:    2,
			t:    3,
			want: false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, ContainsNearbyAlmostDuplicate(tt.nums, tt.k, tt.t))
	}
}
