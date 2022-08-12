package main

import "fmt"

func moveZeroes(nums []int) []int {

	leftPointer := 0
	rightPointer := len(nums) - 1

	for leftPointer < rightPointer {
		if nums[leftPointer] == 0 {
			for i := leftPointer; i < rightPointer; i++ {
				nums[i] = nums[i+1]
			}
			nums[rightPointer] = 0
			rightPointer--
		} else {
			leftPointer++
		}
	}

	return nums
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	fmt.Println(moveZeroes(nums))
}
