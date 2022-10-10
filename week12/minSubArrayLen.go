package main

import (
	"math"
)

//	func minSubArrayLen(target int, nums []int) int {
//		i, n, res := 0, len(nums), len(nums)+1
//		for j := 0; j < n; j++ {
//			target -= nums[j]
//			for target <= 0 {
//				if res > j-i+1 {
//					res = j - i + 1
//				}
//				target += nums[i]
//				i++
//			}
//		}
//		return res % (n + 1)
//	}
func minSubArrayLen(target int, nums []int) int {
	n, ans, left, sum := len(nums), math.MaxInt, 0, 0

	for i := 0; i < n; i++ {
		sum += nums[i]
		for sum >= target {
			if ans > i-i+1 {
				ans = i - left + 1
			}
			sum -= nums[left]
			left++
		}
	}

	if ans != math.MaxInt {
		return ans
	}
	return 0
}

// func main() {
// 	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
// }
