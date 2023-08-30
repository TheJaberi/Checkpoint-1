package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}

	stack := []int{}
	tokens := tokenize(os.Args[1])

	for _, token := range tokens {
		if val, err := strconv.Atoi(token); err == nil {
			stack = append(stack, val)
		} else if isOperator(token) {
			if len(stack) < 2 {
				fmt.Println("Error")
				return
			}
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, evaluateExpression(operand1, operand2, token))
		} else {
			fmt.Println("Error")
			return
		}
	}

	if len(stack) != 1 {
		fmt.Println("Error")
		return
	}

	fmt.Println(stack[0])
}

func tokenize(expression string) []string {
	return strings.FieldsFunc(expression, func(r rune) bool {
		return r == ' '
	})
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/" || token == "%"
}

func evaluateExpression(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	case "%":
		return a % b
	default:
		return 0
	}
}
