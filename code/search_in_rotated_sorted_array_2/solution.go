package search_in_rotated_sorted_array_2

func Search(nums []int, target int) bool {
	return recursivelySearch(nums, target)
}

func recursivelySearch(nums []int, target int) bool {
	n := len(nums)
	if n < 1 {
		return false
	}

	midIdx := n / 2
	start, mid := nums[0], nums[midIdx]

	switch {
	case target == start || target == mid:
		return true
	case start == mid:
		return recursivelySearch(nums[1:], target)
	case start < mid:
		if target > start && target < mid {
			return recursivelySearch(nums[:midIdx], target)
		} else {
			return recursivelySearch(nums[midIdx:], target)
		}
	default:
		if target > mid && target < start {
			return recursivelySearch(nums[midIdx:], target)
		} else {
			return recursivelySearch(nums[:midIdx], target)
		}
	}
}
