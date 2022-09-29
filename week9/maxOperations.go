package main

import (
	"fmt"
	"math"
)

func maxOperations(nums []int, k int) int {
	numMap := make(map[int]int)
	totalOperation := 0
	for _, x := range nums {
		_, ok := numMap[x]
		if ok {
			numMap[x]++
		} else {
			numMap[x] = 1
		}
	}
	for ky, v := range numMap {
		if float64(ky) == float64(k)/2 {
			totalOperation += int(math.Floor(float64(v / 2)))
		} else {
			kminus := k - ky
			if _, ok := numMap[kminus]; ok {
				totalOperation += int(math.Min(float64(v), float64(numMap[kminus])))
				delete(numMap, kminus)
			}
		}
	}
	return totalOperation
}

func main() {
	fmt.Println(maxOperations([]int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2}, 3))
}
