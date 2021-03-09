package best_time_to_buy_and_sell_stock_3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			want:   6,
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			prices: []int{1},
			want:   0,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MaxProfit(tt.prices))
	}
}
