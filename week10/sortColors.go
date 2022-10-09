package main

import "fmt"

func sortColors(nums []int) {
	count := make([]int, 3)
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}
	lastOffSet := 0
	for i := 0; i <= 2; i++ {
		j := lastOffSet
		for ; j < count[i]+lastOffSet; j++ {
			nums[j] = i
		}
		lastOffSet += count[i]

	}

}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
