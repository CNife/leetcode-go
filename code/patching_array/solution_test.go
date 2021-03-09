package patching_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinPatches(t *testing.T) {
	tests := []struct {
		nums []int
		n    int
		want int
	}{
		{
			nums: []int{1, 3},
			n:    6,
			want: 1,
		},
		{
			nums: []int{1, 5, 10},
			n:    20,
			want: 2,
		},
		{
			nums: []int{1, 2, 2},
			n:    5,
			want: 0,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MinPatches(tt.nums, tt.n))
	}
}
