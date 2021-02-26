package min_in_spiral_array

func MinArray(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		switch {
		case nums[mid] < nums[hi]:
			hi = mid
		case nums[mid] > nums[hi]:
			lo = mid + 1
		default:
			hi--
		}
	}
	return nums[lo]
}
