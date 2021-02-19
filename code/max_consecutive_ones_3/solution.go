package max_consecutive_ones_3

// 1004. 最大连续1的个数 III，中等
func LongestOnes(nums []int, k int) int {
	left, zeroCount := 0, 0
	result := 0
	for right := range nums {
		if nums[right] == 0 {
			zeroCount++
		}
		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
