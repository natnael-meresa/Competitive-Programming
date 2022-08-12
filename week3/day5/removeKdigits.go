package main

import (
	"fmt"
	"strings"
)

func removeKdigits(num string, k int) string {
	stk := []byte{}
	removed := 0
	for i := 0; i < len(num); i++ {
		for len(stk) != 0 && removed < k && stk[len(stk)-1] > num[i] {
			stk = stk[:len(stk)-1]
			removed++
		}
		stk = append(stk, num[i])
	}
	stk = stk[:len(stk)-(k-removed)]

	final := strings.TrimLeft(string(stk), "0")
	if final == "" {
		return "0"
	}

	return final
}

func main() {

	fmt.Println(removeKdigits("10", 1))
	// 3432219
	// if small and not much lift
}
