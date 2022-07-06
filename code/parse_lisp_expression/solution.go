package parse_lisp_expression

import (
	"strconv"
	"strings"
)

func Evaluate(expression string) int {
	return evalExpression(expression, nil)
}

func evalExpression(expression string, vars map[string]int) int {
	expression = strings.TrimSpace(expression)
	// 表达式
	if strings.HasPrefix(expression, "(") && strings.HasSuffix(expression, ")") {
		expression = expression[1 : len(expression)-1]
		parts := splitExpression(expression)
		switch parts[0] {
		case "let":
			return evalLet(parts[1:], vars)
		case "add":
			return evalAdd(parts[1:], vars)
		case "mult":
			return evalMult(parts[1:], vars)
		default:
			panic("invalid operation")
		}
	}
	// 数字
	if num, err := strconv.Atoi(expression); err == nil {
		return num
	}
	// 变量
	if value, exists := vars[expression]; exists {
		return value
	} else {
		panic("variable not defined")
	}
}

func copyMap(m map[string]int) map[string]int {
	newMap := make(map[string]int, len(m))
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}

func evalLet(parts []string, vars map[string]int) int {
	thisVars := copyMap(vars)
	for i := 0; i < len(parts)-1; i += 2 {
		varName := parts[i]
		varValue := evalExpression(parts[i+1], thisVars)
		thisVars[varName] = varValue
	}
	return evalExpression(parts[len(parts)-1], thisVars)
}

func evalAdd(parts []string, m map[string]int) int {
	lhs := evalExpression(parts[0], m)
	rhs := evalExpression(parts[1], m)
	return lhs + rhs
}

func evalMult(parts []string, m map[string]int) int {
	lhs := evalExpression(parts[0], m)
	rhs := evalExpression(parts[1], m)
	return lhs * rhs
}

func splitExpression(expression string) []string {
	expression = strings.TrimSpace(expression)

	var result []string
	partStart := 0
	isReadingSubExpression := false
	leftCount := 0
	for i := 0; i < len(expression); i++ {
		switch expression[i] {
		case '(':
			isReadingSubExpression = true
			leftCount++
		case ')':
			leftCount--
			if leftCount < 1 {
				isReadingSubExpression = false
			}
		case ' ':
			if !isReadingSubExpression && partStart < i {
				part := expression[partStart:i]
				result = append(result, part)
				partStart = i + 1
			}
		}
	}
	if !isReadingSubExpression && partStart < len(expression) {
		part := expression[partStart:]
		result = append(result, part)
	}
	return result
}
