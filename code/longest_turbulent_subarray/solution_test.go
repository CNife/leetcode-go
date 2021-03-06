package longest_turbulent_subarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxTurbulenceSize(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{9, 4, 2, 10, 7, 8, 8, 1, 9},
			want: 5,
		},
		{
			nums: []int{4, 8, 12, 16},
			want: 2,
		},
		{
			nums: []int{100},
			want: 1,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MaxTurbulenceSize(tt.nums))
	}
}
