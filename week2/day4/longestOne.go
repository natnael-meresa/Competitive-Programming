package main

import "fmt"

func longestOnes(nums []int, k int) int {
	max, fl, current, startingIndex := 0, 0, 0, 0
	for _, x := range nums {
		if x == 0 && fl >= k {
			if k == 0 {
				fl = 0
				current = 0
			} else {
				for nums[startingIndex] != 0 {
					current -= 1
					startingIndex++
				}
				startingIndex++
			}
		} else if x == 0 && fl < k {
			fl++
			current++
		} else if x == 1 {
			current++
		}

		if current > max {
			max = current
		}
	}

	return max
}

func main() {

	max := longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)
	// max := longestOnes([]int{0, 0, 1, 1, 1, 0, 0}, 0)
	// max := longestOnes([]int{1, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 1}, 9)

	fmt.Println(max)
}
