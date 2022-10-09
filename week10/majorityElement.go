package main

// it is not going to work b/c of -ve numbers
func majorityElement(nums []int) int {
	max, min := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}

	count := make([]int, max-min+1)
	maxNum := nums[0]
	half := len(nums) / 2
	for i := 0; i < len(nums); i++ {
		count[nums[i]-min]++
		if count[nums[i]-min] > half {
			maxNum = nums[i]
			break
		}
	}
	return maxNum
}

// func main() {
// 	fmt.Println(majorityElement([]int{-1, 1, 1, 1, 2, 1}))
// }
