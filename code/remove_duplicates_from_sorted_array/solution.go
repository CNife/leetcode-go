package remove_duplicates_from_sorted_array

func RemoveDuplicates(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	i := 0
	for j := 1; j < n; j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
