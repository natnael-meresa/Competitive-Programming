package main

import (
	"sort"
)

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	maxFrq := 0
	if len(nums) <= 1 {
		return 1
	}

	for i := len(nums) - 1; i >= 1; i-- {
		curFrq, opNum, j := 1, k, i-1
		for opNum > 0 && j >= 0 {
			opNum = nums[j] + opNum - nums[i]
			if opNum < 0 {
				break
			}
			curFrq++
			j--
		}
		if curFrq > maxFrq {
			maxFrq = curFrq
		}
		if j < 0 || maxFrq > i {
			break
		}

	}
	return maxFrq
}

// func main() {
// 	// fmt.Println(maxFrequency([]int{9940, 9995, 9944, 9937, 9941, 9952, 9907, 9952, 9987, 9964, 9940, 9914, 9941, 9933, 9912, 9934, 9980, 9907, 9980, 9944, 9910, 9997}, 7925))
// 	// fmt.Println(maxFrequency([]int{1, 4, 8, 13}, 5))
// 	// fmt.Println(maxFrequency([]int{1, 2, 5}, 5))
// 	fmt.Println(maxFrequency([]int{2, 5}, 5))

// }
