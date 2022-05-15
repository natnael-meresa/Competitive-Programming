package main

import (
	"fmt"
	"sort"
)

func maxCoins(piles []int) int {
	sort.Ints(piles)
	maxCoin, n := 0, len(piles)
	for i, j := 0, n-2; i < (n / 3); i, j = i+1, j-2 {
		maxCoin += piles[j]
	}
	return maxCoin
}

func main() {
	maxCoin := maxCoins([]int{9, 8, 7, 6, 5, 1, 2, 3, 4})
	fmt.Println(maxCoin)
}
