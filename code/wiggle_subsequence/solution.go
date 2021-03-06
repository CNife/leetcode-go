package wiggle_subsequence

func WiggleMaxLength(nums []int) int {
	up, down := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	switch {
	case len(nums) == 0:
		return 0
	case up > down:
		return up
	default:
		return down
	}
}
