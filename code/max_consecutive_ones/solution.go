package max_consecutive_ones

func FindMaxConsecutiveOnes(nums []int) int {
	count, maxCount := 0, 0
	for _, num := range nums {
		if num == 1 {
			count++
		} else if count > 0 {
			if maxCount < count {
				maxCount = count
			}
			count = 0
		}
	}
	if maxCount < count {
		return count
	}
	return maxCount
}
