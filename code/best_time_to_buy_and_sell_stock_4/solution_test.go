package best_time_to_buy_and_sell_stock_4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		k      int
		prices []int
		want   int
	}{
		{
			k:      2,
			prices: []int{2, 4},
			want:   2,
		},
		{
			k:      2,
			prices: []int{3, 2, 6, 5, 0, 3},
			want:   7,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MaxProfit(tt.k, tt.prices))
	}
}
