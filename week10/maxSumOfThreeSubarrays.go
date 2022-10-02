package main

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	W := make([]int, len(nums)-1)

	currSum := 0
	for i := 0; i < k; i++ {
		currSum += nums[i]
	}
	W[0] = currSum

	for i := k; i < len(nums); i++ {
		currSum -= nums[i-k]
		currSum += nums[i]
		W[i-k+1] = currSum
	}
	wlen := len(W)
	left := make([]int, wlen)
	best := 0
	for i := 0; i < wlen; i++ {
		if W[i] > W[best] {
			best = i
		}
		left[i] = best
	}

	right := make([]int, wlen)
	best = wlen - 1
	for i := wlen - 1; i >= 0; i-- {
		if W[i] >= W[best] {
			best = i
		}
		right[i] = best
	}

	ans := []int{-1, -1, -1}
	for j := k; j < wlen-k; j++ {
		i, l := left[j-k], right[j+k]
		if ans[0] == -1 || W[i]+W[j]+W[l] > W[ans[0]]+W[ans[1]]+W[ans[2]] {
			ans[0] = i
			ans[1] = j
			ans[2] = l
		}
	}
	return ans
}
