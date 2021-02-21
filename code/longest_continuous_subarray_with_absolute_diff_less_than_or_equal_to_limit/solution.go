package longest_continuous_subarray_with_absolute_diff_less_than_or_equal_to_limit

// 1438. 绝对差不超过限制的最长连续子数组，中等
func LongestSubarray(nums []int, limit int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	maxQueue, minQueue := queue{}, queue{}
	left, right, result := 0, 0, 0

	for right < n {
		for len(maxQueue) > 0 && nums[right] > maxQueue.peekRight() {
			maxQueue.popRight()
		}
		for len(minQueue) > 0 && nums[right] < minQueue.peekRight() {
			minQueue.popRight()
		}

		maxQueue.pushRight(nums[right])
		minQueue.pushRight(nums[right])
		right++

		for maxQueue.peekLeft()-minQueue.peekLeft() > limit {
			if maxQueue.peekLeft() == nums[left] {
				maxQueue.popLeft()
			}
			if minQueue.peekLeft() == nums[left] {
				minQueue.popLeft()
			}
			left++
		}

		if length := right - left; result < length {
			result = length
		}
	}

	return result
}

type queue []int

func (q queue) peekLeft() int {
	return q[0]
}

func (q queue) peekRight() int {
	return q[len(q)-1]
}

func (q *queue) popLeft() int {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func (q *queue) popRight() int {
	n := len(*q) - 1
	x := (*q)[n]
	*q = (*q)[:n]
	return x
}

func (q *queue) pushRight(x int) {
	*q = append(*q, x)
}
