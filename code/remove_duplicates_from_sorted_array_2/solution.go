package remove_duplicates_from_sorted_array_2

func RemoveDuplicates(nums []int) int {
	n := len(nums)
	if n < 3 {
		return n
	}

	left, count := 1, 1
	for right := 1; right < n; right++ {
		if nums[right-1] == nums[right] {
			count++
		} else {
			count = 1
		}
		if count <= 2 {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}
