package best_time_to_buy_and_sell_stock_3

func MaxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	var firstSold, secondSold int
	firstBought, secondBought := -prices[0], -prices[0]

	for i := 1; i < n; i++ {
		price := prices[i]
		secondSold = max(secondSold, secondBought+price)
		secondBought = max(secondBought, firstSold-price)
		firstSold = max(firstSold, firstBought+price)
		firstBought = max(firstBought, -price)
	}

	return secondSold
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
