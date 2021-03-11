package basic_calculator_2

func Calculate(s string) int {
	var numStack []int
	var opStack []byte

	for {
		t, next, success := readNextToken(s)
		if !success {
			break
		}

		switch t.tag {
		case number:
			numStack = append(numStack, t.number)
		case operator:
			for len(opStack) > 0 {
				top := opStack[len(opStack)-1]
				if priority(top) >= priority(t.operator) {
					opStack = opStack[:len(opStack)-1]
					calc(&numStack, top)
				} else {
					break
				}
			}
			opStack = append(opStack, t.operator)
		default:
			panic("invalid token tag")
		}

		s = next
	}

	for len(opStack) > 0 {
		n := len(opStack) - 1
		op := opStack[n]
		opStack = opStack[:n]
		calc(&numStack, op)
	}

	//goland:noinspection GoNilness
	return numStack[len(numStack)-1]
}

type token struct {
	tag      tokenTag
	operator byte
	number   int
}

type tokenTag byte

const (
	number tokenTag = iota + 1
	operator
)

func readNextToken(input string) (token, string, bool) {
	i := 0
	for ; i < len(input); i++ {
		if input[i] != ' ' {
			break
		}
	}
	if i >= len(input) {
		return token{}, "", false
	}

	first := input[i]
	if first >= '0' && first <= '9' {
		num, j := 0, i
		for ; j < len(input); j++ {
			if ch := input[j]; ch >= '0' && ch <= '9' {
				num = 10*num + int(ch-'0')
			} else {
				break
			}
		}
		return token{tag: number, number: num}, input[j:], true
	}
	return token{tag: operator, operator: first}, input[i+1:], true
}

func calc(numStack *[]int, op byte) {
	n := len(*numStack)
	rhs := (*numStack)[n-1]
	top := &(*numStack)[n-2]
	switch op {
	case '+':
		*top += rhs
	case '-':
		*top -= rhs
	case '*':
		*top *= rhs
	case '/':
		*top /= rhs
	default:
		panic("invalid operator")
	}
	*numStack = (*numStack)[:n-1]
}

func priority(op byte) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	default:
		panic("invalid operator")
	}
}
