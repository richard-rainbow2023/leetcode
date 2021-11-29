package letture02

import (
	"container/list"
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := list.New()

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			num1 := stack.Back().Value.(int)
			stack.Remove(stack.Back())
			num2 := stack.Back().Value.(int)
			stack.Remove(stack.Back())
			result := calculate(num2, num1, token)
			stack.PushBack(result)
		default:
			num, _ := strconv.Atoi(token)
			stack.PushBack(num)
		}
	}
	return stack.Back().Value.(int)
}

func calculate(num1 int, num2 int, operator string) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		fmt.Println("/ is", num1/num2)
		return num1 / num2
	default:
		return 0
	}
}
