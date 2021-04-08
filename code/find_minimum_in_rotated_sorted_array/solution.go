package find_minimum_in_rotated_sorted_array

func FindMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		middle := left + (right-left)/2
		if nums[middle] > nums[right] {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return nums[left]
}

// func FindMin(nums []int) int {
// 	n := len(nums)
// 	switch n {
// 	case 0:
// 		panic("empty array")
// 	case 1:
// 		return nums[0]
// 	case 2:
// 		return min(nums[0], nums[1])
// 	}
//
// 	first, last := nums[0], nums[n-1]
// 	if first < last {
// 		return first
// 	}
// 	middleIndex := n / 2
// 	middle := nums[middleIndex]
// 	if middle < last {
// 		return FindMin(nums[:middleIndex+1])
// 	} else {
// 		return FindMin(nums[middleIndex:])
// 	}
// }
//
// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }
