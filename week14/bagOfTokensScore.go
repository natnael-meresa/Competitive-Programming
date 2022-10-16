package main

import (
	"fmt"
	"sort"
)

func bagOfTokensScore(tokens []int, power int) int {
	n := len(tokens)
	score, l, r := 0, 0, n-1

	sort.Ints(tokens)

	for l < r {
		if power < tokens[l] {
			if score >= 1 && r-l > 1 {
				power += tokens[r]
				score--
			} else {
				break
			}

			r--
		} else {
			power -= tokens[l]
			score++
			l++
		}
	}
	if score < 1 {
		return 0
	} else {
		return score
	}

}

func main() {
	fmt.Println(bagOfTokensScore([]int{100, 200}, 150))
}
