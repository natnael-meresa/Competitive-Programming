package main

import "fmt"

func isValid(s string) bool {
	stack := []string{}

	for _, x := range s {
		if string(x) == "(" || string(x) == "{" || string(x) == "[" {
			stack = append(stack, string(x))
		} else {

			n := len(stack) - 1
			if n+1 == 0 {
				stack = append(stack, string(x))
			} else if (string(stack[n]) == "(" && string(x) == ")") || (string(stack[n]) == "[" && string(x) == "]") || (string(stack[n]) == "{" && string(x) == "}") {
				stack = stack[:n]
			} else {
				stack = append(stack, string(x))
			}
		}
	}

	return len(stack) == 0
}

func main() {
	valid := isValid("{}")

	fmt.Println(valid)
}
