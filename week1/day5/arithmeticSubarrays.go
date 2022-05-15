package main

import "sort"

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	result := make([]bool, len(l))
	tempNums := make([]int, len(nums))

	for i := 0; i < len(l); i++ {
		copy(tempNums, nums)

		temp := tempNums[l[i] : r[i]+1]
		sort.Ints(temp)
		result[i] = true

		for j := 0; j < len(temp)-1; j++ {
			if temp[j+1]-temp[j] != temp[1]-temp[0] {
				result[i] = false
				break
			}
		}
	}
	return result
}

func main() {

	nums := []int{4, 6, 5, 9, 3, 7}
	l := []int{0, 0, 2}
	r := []int{2, 3, 5}
	checkArithmeticSubarrays(nums, l, r)
}
