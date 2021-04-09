package find_minimum_in_rotated_sorted_array_2

func FindMin(nums []int) int {
	n := len(nums)
	l, r := 0, n-1
	for l < r && nums[0] == nums[r] {
		r--
	}
	for l < r {
		m := (l + r + 1) >> 1
		if nums[m] >= nums[0] {
			l = m
		} else {
			r = m - 1
		}
	}
	if r+1 < n {
		return nums[r+1]
	} else {
		return nums[0]
	}
}
