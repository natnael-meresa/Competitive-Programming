package main

func rotate(nums []int, k int) {
	tempNums := make([][]int, len(nums))
	ix, it := 0, 0

	for i := 0; i < len(nums); i++ {
		tempI := i + k
		for tempI > len(nums)-1 {
			tempI = tempI - len(nums)
		}

		tempNums[it] = []int{nums[tempI], tempI}
		it++
		if tempNums[ix][1] == i {
			nums[tempI] = tempNums[ix][0]
			ix++
		} else {
			nums[tempI] = nums[i]
		}
	}
}

// func main() {
// 	bRotate := []int{1, 2}
// 	rotate(bRotate, 9)
// 	fmt.Println(bRotate)
// }

// // 1,2
