package main

import "fmt"

func maxArea(height []int) int {
	start, end := 0, len(height)-1
	max_area := 0
	for start < end {
		area := min(height[start], height[end]) * (end - start)

		if height[start] >= height[end] {
			end--
		} else {
			start++
		}
		max_area = max(area, max_area)
	}

	return max_area
}

func min(x, y int) int {
	if y < x {
		return y
	}

	return x
}

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
