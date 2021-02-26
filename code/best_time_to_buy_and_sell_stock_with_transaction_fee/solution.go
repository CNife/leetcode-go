package best_time_to_buy_and_sell_stock_with_transaction_fee

func MaxProfit(prices []int, fee int) int {
	m, n := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		m, n = max(m, n-prices[i]), max(n, m+prices[i]-fee)
	}
	return max(m, n)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
