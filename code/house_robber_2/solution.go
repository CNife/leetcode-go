package house_robber_2

func Rob(nums []int) int {
	helper := func(slice []int) int {
		prev, curr := 0, 0
		for _, num := range slice {
			prev, curr = curr, max(curr, prev+num)
		}
		return curr
	}

	switch n := len(nums); n {
	case 0:
		return 0
	case 1:
		return nums[0]
	default:
		return max(helper(nums[1:]), helper(nums[:n-1]))
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
