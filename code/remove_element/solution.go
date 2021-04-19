package remove_element

func RemoveElement(nums []int, val int) int {
	i, n := 0, len(nums)
	for j := 0; j < n; j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
