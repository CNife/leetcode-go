package volume_of_histogram_lcci

func Trap(height []int) int {
	heightSum, volume := sum(height), 0
	for left, right, high := 0, len(height)-1, 1; left <= right; high++ {
		for left <= right && height[left] < high {
			left++
		}
		for left <= right && height[right] < high {
			right--
		}
		volume += right - left + 1
	}
	return volume - heightSum
}

func sum(nums []int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}
