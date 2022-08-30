package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		num1 := strconv.Itoa(nums[i])
		num2 := strconv.Itoa(nums[j])

		return num1+num2 >= num2+num1
	})

	if nums[0] == 0 {
		return "0"
	}

	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nums)), ""), "[]")
}

func main() {
	large := largestNumber([]int{3, 30, 34, 5, 9})
	fmt.Println(large)
}
