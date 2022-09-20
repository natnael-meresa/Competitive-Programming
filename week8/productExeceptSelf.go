package main

func productExceptSelf(nums []int) []int {
	prefixSum := make([]int, len(nums))
	suffixSum := make([]int, len(nums))
	answer := make([]int, len(nums))

	prefixSum[0], suffixSum[len(nums)-1] = nums[0], nums[len(nums)-1]
	for i, j := 1, len(nums)-2; j >= 0; i, j = i+1, j-1 {
		prefixSum[i] = prefixSum[i-1] * nums[i]
		suffixSum[j] = suffixSum[j+1] * nums[j]
	}

	answer[0] = suffixSum[1]
	answer[len(nums)-1] = prefixSum[len(nums)-2]

	for i := 1; i < len(nums)-1; i++ {
		answer[i] = prefixSum[i-1] * suffixSum[i+1]
	}
	return answer
}

// func main() {
// 	fmt.Println(">>", productExceptSelf([]int{1, 2, 3, 4}))
// }
