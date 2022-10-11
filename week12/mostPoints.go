package main

import (
	"fmt"
)

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	preSum := make([]int, n)
	max := 0
	for i := 0; i < n; i++ {
		if i > 0 {
			if preSum[i-1] > preSum[i] {
				preSum[i] = preSum[i-1]
			}
		}

		next := i + questions[i][1] + 1
		if next < n {
			tempSum := preSum[i] + questions[i][0]
			if tempSum > preSum[next] {
				preSum[next] = tempSum
			}
		}
	}

	for i := 0; i < len(preSum); i++ {
		if curPoints := preSum[i] + questions[i][0]; curPoints > max {
			max = curPoints
		}
	}

	return int64(max)
}

func main() {
	fmt.Println(mostPoints([][]int{{21, 5}, {92, 3}, {74, 2}, {39, 4}, {58, 2}, {5, 5}, {49, 4}, {65, 3}}))
}
