package main

import "fmt"

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	maxCount := countMax(nums, firstLen, secondLen)
	secondMax := countMax(nums, secondLen, firstLen)
	fmt.Println(maxCount, secondMax)
	if secondMax > maxCount {
		return secondMax
	}
	return maxCount

}

func countMax(nums []int, firstLen int, secondLen int) int {
	maxCount := 0
	currentFirst := 0
	currentSecond := 0

	for i := len(nums) - (secondLen); i >= firstLen; i-- {

		for k := i; k < i+secondLen; k++ {
			currentSecond += nums[k]
		}

		for j := firstLen; j <= i; j++ {

			for x := firstLen; x > 0; x-- {
				currentFirst += nums[j-x]
			}

			if currentFirst+currentSecond > maxCount {
				maxCount = currentFirst + currentSecond
			}
			currentFirst = 0
		}
		currentSecond = 0
	}
	return maxCount
}

func main() {
	fmt.Println(maxSumTwoNoOverlap([]int{1, 0, 3}, 1, 2))
}
