package main

import "fmt"

type KthLargest struct {
	k    int
	nums []int
}

func downHeapify(nums []int, i int) {
	lowset := i
	left := 2*i + 1
	right := 2*i + 2

	if left < len(nums) && nums[left] < nums[lowset] {
		lowset = left
	}

	if right < len(nums) && nums[right] < nums[lowset] {
		lowset = right
	}

	if lowset != i {
		nums[lowset], nums[i] = nums[i], nums[lowset]
		downHeapify(nums, lowset)
	}
}
func pop(nums *[]int) {
	l := len(*nums) - 1
	(*nums)[0] = (*nums)[l]
	*nums = (*nums)[:l]
	downHeapify((*nums), 0)
}

func heapifyUP(nums []int, i int) {
	for nums[i] < nums[pIndex(i)] {
		nums[pIndex(i)], nums[i] = nums[i], nums[pIndex(i)]
		i = pIndex(i)
	}
}

func Constructor(k int, nums []int) KthLargest {
	startIndex := (len(nums) / 2) - 1

	for i := startIndex; i >= 0; i-- {
		downHeapify(nums, i)
	}
	for len(nums) > k {
		pop(&nums)
	}

	return KthLargest{
		k:    k,
		nums: nums,
	}
}

func pIndex(i int) int {
	return (i - 1) / 2
}
func (this *KthLargest) Add(val int) int {
	this.nums = append(this.nums, val)
	heapifyUP(this.nums, len(this.nums)-1)

	for len(this.nums) > this.k {
		pop(&this.nums)
	}
	return this.nums[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func main() {
	obj := Constructor(3, []int{4, 5, 8, 2})
	param_1 := obj.Add(3)

	fmt.Println(param_1)
}
