package main

import "fmt"

func partitionLabels(s string) []int {
	last := make(map[byte]int, 26)
	ans := make([]int, 0)
	for i := 0; i < len(s); i++ {
		last[s[i]] = i
	}

	j, a := 0, 0
	for i := 0; i < len(s); i++ {
		if last[s[i]] > j {
			j = last[s[i]]
		}

		if j == i {
			ans = append(ans, i-a+1)
			a = i + 1
		}
	}

	return ans
}

func main() {
	fmt.Println(partitionLabels("eccbbbbdec"))
}
