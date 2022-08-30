package main

import (
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Slice(people, func(i, j int) bool {
		return people[i] <= people[j]
	})

	l, h := 0, len(people)-1
	ans := 0
	for l <= h {
		ans++
		if people[l]+people[h] <= limit {
			l++
		}
		h--
	}
	return ans
}

// func main() {
// 	fmt.Println(numRescueBoats([]int{7, 2, 3, 5, 2}, 5))
// }
