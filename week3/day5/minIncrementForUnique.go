package main

import (
	"fmt"
	"sort"
)

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			count += nums[i-1] - nums[i] + 1
			nums[i] = nums[i-1] + 1
		}
	}
	return count
}

func main() {
	fmt.Println(minIncrementForUnique([]int{3, 2, 1, 2, 1, 7}))
	// 1,1,2,2,3,7
}

// // https://leetcode.com/problems/minimum-increment-to-make-array-unique/solution/
