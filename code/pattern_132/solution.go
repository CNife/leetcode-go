package pattern_132

import "math"

//goland:noinspection GoNilness
func Find132Pattern(nums []int) bool {
	n := len(nums)

	leftMin := make([]int, n)
	for i := 1; i < n; i++ {
		leftMin[i] = min(leftMin[i-1], nums[i-1])
	}

	var stack []int
	for i := n - 1; i >= 0; i-- {
		numsK := math.MinInt32
		for {
			stackLen := len(stack)
			if stackLen > 0 && stack[stackLen-1] < nums[i] {
				numsK = stack[stackLen-1]
				stack = stack[:stackLen-1]
			} else {
				break
			}
		}
		if leftMin[i] < numsK {
			return true
		}
		stack = append(stack, nums[i])
	}
	return false
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
