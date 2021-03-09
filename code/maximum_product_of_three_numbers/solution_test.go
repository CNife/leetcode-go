package maximum_product_of_three_numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumProduct(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 2, 3},
			want: 6,
		},
		{
			nums: []int{1, 2, 3, 4},
			want: 24,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MaximumProduct(tt.nums))
	}
}
