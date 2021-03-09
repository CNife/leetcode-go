package best_time_to_buy_and_sell_stock_with_transaction_fee

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices    []int
		fee, want int
	}{
		{
			prices: []int{1, 3, 2, 8, 4, 9},
			fee:    2,
			want:   8,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MaxProfit(tt.prices, tt.fee))
	}
}
