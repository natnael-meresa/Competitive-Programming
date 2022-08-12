package main

func subarraySum(nums []int, k int) int {
	cur_sum, num := 0, 0
	freq := map[int]int{0: 1}

	for i := 0; i < len(nums); i++ {
		cur_sum += nums[i]
		num += freq[cur_sum-k]
		freq[cur_sum]++
	}

	return num
}
