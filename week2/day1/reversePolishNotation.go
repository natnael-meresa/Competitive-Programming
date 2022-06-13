package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	Items []string
}

func Constructor() Stack {
	stack := Stack{}
	return stack
}

func (s *Stack) Push(x string) {
	s.Items = append(s.Items, x)
}

func (s *Stack) Pop() int {
	n := len(s.Items) - 1

	last, _ := strconv.Atoi(s.Items[n])
	s.Items = s.Items[:n]
	return last
}

func evalRPN(tokens []string) int {
	oprand1, oprand2, tempResult, stack := 0, 0, 0, Stack{}
	for _, x := range tokens {

		switch string(x) {
		case "+", "-", "*", "/":
			oprand1 = stack.Pop()
			oprand2 = stack.Pop()
			switch string(x) {
			case "*":
				tempResult = oprand1 * oprand2
			case "/":
				tempResult = oprand2 / oprand1
			case "+":
				tempResult = oprand1 + oprand2
			case "-":
				tempResult = oprand2 - oprand1
			}
			stack.Push(strconv.Itoa(tempResult))
		default:
			stack.Push(x)
		}
	}
	final, _ := strconv.Atoi(stack.Items[0])
	return final

}

func main() {
	result := evalRPN([]string{"4", "13", "5", "-", "+"})

	fmt.Println(result)

}
