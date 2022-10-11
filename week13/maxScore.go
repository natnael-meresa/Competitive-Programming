package main

func maxScore(cardPoints []int, k int) int {

	curTotal := 0

	for i := 0; i < k; i++ {
		curTotal += cardPoints[i]
	}

	maxTotal := curTotal
	for i := 0; i < k; i++ {
		curTotal -= cardPoints[k-i-1]
		curTotal += cardPoints[len(cardPoints)-i-1]
		if curTotal > maxTotal {
			maxTotal = curTotal
		}
	}

	return maxTotal
}
