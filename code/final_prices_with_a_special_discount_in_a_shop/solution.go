package final_prices_with_a_special_discount_in_a_shop

func FinalPrices(prices []int) []int {
	result := make([]int, len(prices))
	copy(result, prices)
	for i, price := range prices {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= price {
				result[i] -= prices[j]
				break
			}
		}
	}
	return result
}
