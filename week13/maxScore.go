package main

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	curTotal := 0

	for i := 0; i < k; i++ {
		curTotal += cardPoints[i]
	}

	maxTotal := curTotal
	for i := 1; i <= k; i++ {
		curTotal -= cardPoints[k-i]
		curTotal += cardPoints[n-i]
		if curTotal > maxTotal {
			maxTotal = curTotal
		}
	}

	return maxTotal
}
