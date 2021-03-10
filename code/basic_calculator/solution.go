package basic_calculator

//goland:noinspection GoNilness
func Calculate(s string) int {
	result, num, sign := 0, 0, 1
	var stack []int
	for i := 0; i < len(s); i++ {
		b := s[i]
		switch {
		case b >= '0' && b <= '9':
			num = 10*num + int(b-'0')
		case b == '+' || b == '-':
			result += sign * num
			num = 0
			if b == '+' {
				sign = 1
			} else {
				sign = -1
			}
		case b == '(':
			stack = append(stack, result, sign)
			result, sign = 0, 1
		case b == ')':
			result += sign * num
			num = 0
			stackLen := len(stack)
			result *= stack[stackLen-1]
			result += stack[stackLen-2]
			stack = stack[:stackLen-2]
		}
	}
	result += sign * num
	return result
}
