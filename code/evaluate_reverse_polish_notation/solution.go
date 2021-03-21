package evaluate_reverse_polish_notation

import (
	"log"
	"strconv"
)

//goland:noinspection GoNilness
func EvalRPN(tokens []string) int {
	var stack []int

	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			stack = append(stack, num)
			continue
		}

		n := len(stack)
		lhs, rhs := &stack[n-2], stack[n-1]
		stack = stack[:n-1]

		switch token[0] {
		case '+':
			*lhs += rhs
		case '-':
			*lhs -= rhs
		case '*':
			*lhs *= rhs
		case '/':
			*lhs /= rhs
		default:
			log.Fatalf("invalid operator %v", token)
		}
	}

	return stack[0]
}
