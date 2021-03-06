package create_maximum_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxNumber(t *testing.T) {
	tests := []struct {
		nums1, nums2 []int
		k            int
		want         []int
	}{
		{
			nums1: []int{3, 4, 6, 5},
			nums2: []int{9, 1, 2, 5, 8, 3},
			k:     5,
			want:  []int{9, 8, 6, 5, 3},
		},
		{
			nums1: []int{6, 7},
			nums2: []int{6, 0, 4},
			k:     5,
			want:  []int{6, 7, 6, 0, 4},
		},
		{
			nums1: []int{3, 9},
			nums2: []int{8, 9},
			k:     3,
			want:  []int{9, 8, 9},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MaxNumber(tt.nums1, tt.nums2, tt.k))
	}
}
