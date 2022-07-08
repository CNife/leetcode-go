package minimum_cost_to_move_chips_to_the_same_position

func MinCostToMoveChips(position []int) int {
	oddCount, evenCount := 0, 0
	for _, num := range position {
		if num%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	if evenCount < oddCount {
		return evenCount
	}
	return oddCount
}
