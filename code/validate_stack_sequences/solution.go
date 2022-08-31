package validate_stack_sequences

func ValidateStackSequences(pushed, popped []int) bool {
	var stack []int
	i := 0
	for _, num := range pushed {
		for i < len(popped) && len(stack) > 0 && popped[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			i++
		}
		stack = append(stack, num)
	}
	for i < len(popped) && len(stack) > 0 && popped[i] == stack[len(stack)-1] {
		stack = stack[:len(stack)-1]
		i++
	}
	return len(stack) == 0
}
