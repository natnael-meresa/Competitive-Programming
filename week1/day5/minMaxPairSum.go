package main

import (
	"fmt"
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)

	fmt.Println(nums)
	max, n := 0, len(nums)
	for i := 0; i < n/2; i++ {
		fmt.Println(n - (i + 1))
		if nums[i]+nums[n-(i+1)] > max {
			max = nums[i] + nums[n-(i+1)]
		}
	}

	return max
}
func main() {
	maxPair := minPairSum([]int{4, 1, 5, 1, 2, 5, 1, 5, 5, 4})

	fmt.Println(maxPair)
}
