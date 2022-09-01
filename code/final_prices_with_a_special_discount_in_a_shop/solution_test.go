package final_prices_with_a_special_discount_in_a_shop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFinalPrices(t *testing.T) {
	tests := []struct {
		prices, want []int
	}{
		{
			prices: []int{8, 4, 6, 2, 3},
			want:   []int{4, 2, 4, 2, 3},
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			prices: []int{10, 1, 1, 6},
			want:   []int{9, 0, 1, 6},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, FinalPrices(tt.prices))
	}
}
