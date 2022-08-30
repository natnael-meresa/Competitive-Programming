package main

import (
	"math"
	"sort"
)

func leastInterval(tasks []byte, n int) int {
	char_count := [26]int{}

	for _, x := range tasks {
		char_count[x-[]byte("A")[0]]++

	}
	sort.IntSlice(char_count[:]).Sort()
	max_count := char_count[25] - 1
	idleSlot := max_count * n

	for i := 24; i >= 0; i-- {
		if char_count[i] == 0 {
			continue
		}
		idleSlot -= int(math.Min(float64(char_count[i]), float64(max_count)))

	}
	if idleSlot < 0 {
		idleSlot = 0
	}
	return len(tasks) + idleSlot
}

// func main() {
// 	fmt.Println(leastInterval([]byte("ABDEYYYYYNNNNNNNNNNNNNNNNN"), 2))
// }
