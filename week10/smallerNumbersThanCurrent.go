package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	output := make([]int, len(nums))

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	count := make([]int, max+1)
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}

	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			output[i] = 0
			continue
		}
		output[i] = count[nums[i]-1]
	}

	return output
}

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
}
