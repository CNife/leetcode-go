package degree_of_an_array

// 697. 数组的度，简单
func FindShortestSubArray(nums []int) int {
	type counterItem struct {
		index, count int
	}

	maxCount, minWindow := 0, 0
	counter := map[int]counterItem{}

	for i, num := range nums {
		var index, count int
		if item, exists := counter[num]; exists {
			index, count = item.index, item.count+1
		} else {
			index, count = i, 1
		}
		counter[num] = counterItem{index, count}

		window := i - index + 1
		if count > maxCount {
			maxCount = count
			minWindow = window
		} else if count == maxCount && minWindow > window {
			minWindow = window
		}
	}

	return minWindow
}
