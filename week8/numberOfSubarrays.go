package main

import "fmt"

// import "fmt"

func numberOfSubarrays(nums []int, k int) int {
	prefixSums := map[int]int{0: 1}
	sum, res := 0, 0
	for _, n := range nums {
		fmt.Println(prefixSums, ">>", res)
		sum += n % 2
		res += prefixSums[sum-k]
		prefixSums[sum]++
	}
	return res
}

// func main() {
// 	fmt.Println(numberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2))
// }
