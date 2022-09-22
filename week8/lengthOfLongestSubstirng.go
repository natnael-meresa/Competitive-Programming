package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	maxCount := 0
	window := map[rune]int{}
	windowStart := 0
	currentCount := 0
	for i, x := range s {
		if s, ok := window[x]; ok {
			if s >= windowStart {
				windowStart = s + 1
			}

			if currentCount > maxCount {
				maxCount = currentCount
			}
			currentCount = i - windowStart
			fmt.Println(">>", currentCount, i, string(x))
			delete(window, x)
		}
		window[x] = i
		currentCount++
		fmt.Println("<<", currentCount, i, string(x))

	}

	if currentCount > maxCount {
		maxCount = currentCount
	}

	return maxCount
}

// func main() {
// 	fmt.Println(lengthOfLongestSubstring("abba"))
// }
