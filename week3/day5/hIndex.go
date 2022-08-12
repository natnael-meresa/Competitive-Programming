package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	// sort.Ints(citations)
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	h := 0
	for i := 0; i < len(citations); i++ {
		if h >= citations[i] {
			break
		}
		h++
	}

	return h
}

func main() {
	fmt.Println(hIndex([]int{1, 3, 1}))
}
