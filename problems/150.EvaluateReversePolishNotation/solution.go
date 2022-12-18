package problems

import (
	. "leetcode-go/adt"

	"strconv"
)

var operators = map[string](func(int, int) int){
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
}

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		applyOperator, isOperator := operators[token]
		if isOperator {
			result := applyOperator(stack[len(stack)-2], stack[len(stack)-1])
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func evalRPN_Stack(tokens []string) int {
	stack := NewStack()

	for _, token := range tokens {
		applyOperator, isOperator := operators[token]
		if isOperator {
			operand1 := stack.Pop()
			operand2 := stack.Pop()
			result := applyOperator(operand2, operand1)
			stack.Push(result)
		} else {
			integer, _ := strconv.Atoi(token)
			stack.Push(integer)
		}
	}

	return stack.Peek()
}
