package next_greater_element_2

func NextGreaterElements(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := range result {
		result[i] = -1
	}

	var s stack
	for i := 0; i < n*2; i++ {
		index := i % n
		for !s.empty() && nums[s.peek()] < nums[index] {
			result[s.pop()] = nums[index]
		}
		s.push(index)
	}

	return result
}

type stack []int

func (s stack) empty() bool {
	return len(s) < 1
}

func (s stack) peek() int {
	return s[len(s)-1]
}

func (s *stack) push(x int) {
	*s = append(*s, x)
}

func (s *stack) pop() int {
	n := len(*s) - 1
	x := (*s)[n]
	*s = (*s)[:n]
	return x
}
