package main

// import (
// 	"fmt"
// 	"math"
// )

// // func PredictTheWinner(nums []int) bool {
// // 	player1, player2 := 0, 0
// // 	turn := true
// // 	localMax := 0
// // 	for len(nums) > 1 {
// // 		if len(nums) > 5 {
// // 			if int(math.Min(float64(nums[len(nums)-1]), float64(nums[1])))+nums[2]+nums[len(nums)-2] > int(math.Max(float64(nums[0]), float64(nums[len(nums)-1])))+nums[1]+nums[len(nums)-3] {
// // 				localMax = nums[0]
// // 				nums = nums[1:]
// // 			} else {
// // 				localMax = nums[len(nums)-1]
// // 				nums = nums[:len(nums)-1]
// // 			}
// // 		} else {
// // 			if nums[0]+nums[len(nums)-2] >= nums[1]+nums[len(nums)-1] {
// // 				localMax = nums[0]
// // 				nums = nums[1:]

// // 			} else {
// // 				localMax = nums[len(nums)-1]
// // 				nums = nums[:len(nums)-1]
// // 			}
// // 		}

// // 		if turn {
// // 			fmt.Println(localMax)
// // 			player1 += localMax
// // 			turn = false
// // 		} else {
// // 			player2 += localMax
// // 			turn = true
// // 		}
// // 	}
// // 	if turn {
// // 		player1 += nums[0]
// // 	} else {
// // 		player2 += nums[0]
// // 	}

// // 	return player1 >= player2
// // }

// func PredictTheWinner(nums []int) bool {

// 	memo := make([][]int, len(nums))
// 	for i := 0; i < len(memo); i++ {
// 		memo[i] = make([]int, len(nums))
// 	}
// 	return winner(nums, 0, len(nums)-1, memo) >= 0
// }
// func winner(nums []int, s, e int, memo [][]int) int {
// 	if s == e {
// 		return nums[s]
// 	}
// 	if memo[s][e] != 0 {
// 		return memo[s][e]
// 	}
// 	a := nums[s] - winner(nums, s+1, e, memo)
// 	b := nums[e] - winner(nums, s, e-1, memo)
// 	memo[s][e] = int(math.Max(float64(a), float64(b)))
// 	return memo[s][e]
// }
// func main() {
// 	fmt.Println(PredictTheWinner([]int{3606449, 6, 5, 9, 452429, 7, 9580316, 9857582, 8514433, 9, 6, 6614512, 753594, 5474165, 4, 2697293, 8, 7, 1}))
// }

// // 3606449,6,5,9,452429,7,9580316,9857582,8514433,9,6,6614512,753594,5474165,4,2697293,8,7,1
// // [3606449,6,5,9,452429,7,9580316,9857582,8514433,9,6,6614512,753594,5474165,4,2697293]
// // [3606449,6,5,5474165,4,2697293]
// // 3606449 , 6, 9 , 4
// // 2697293 , 5, 452429 , 5474165]
