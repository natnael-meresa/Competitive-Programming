package main

import "fmt"

func reverseParentheses(s string) string {
	actual, stack := make([]byte, 0), make([][]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, actual)
			actual = make([]byte, 0)
		} else if s[i] == ')' {

			for i := 0; i < (len(actual))/2; i++ {
				actual[i], actual[len(actual)-1-i] = actual[len(actual)-1-i], actual[i]
			}
			if len(stack) > 0 {
				actual = append(stack[len(stack)-1], actual...)
				stack = stack[:len(stack)-1]
			}
		} else {
			actual = append(actual, s[i])
		}
	}
	return string(actual)
}

func main() {

	hey := reverseParentheses("(ed(et(oc))el)")

	fmt.Println(hey)
}
