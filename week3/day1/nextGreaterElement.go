package main

import "fmt"

type Stack struct {
	items []int
}

func Constructor() Stack {
	item := Stack{}
	return item
}

func (st *Stack) Push(val int) {
	st.items = append(st.items, val)
}

func (st *Stack) Pop() {
	st.items = st.items[:len(st.items)-1]

}

func (st *Stack) Top() int {
	return st.items[len(st.items)-1]
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	answer := make([]int, len(nums1))
	indexes := make(map[int]int, 0)
	for i, n := range nums1 {
		indexes[n] = i
		answer[i] = -1
	}
	stack := Constructor()
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack.items) != 0 && stack.Top() <= nums2[i] {
			fmt.Println(stack)
			stack.Pop()
		}
		if i, ok := indexes[nums2[i]]; ok {
			if len(stack.items) > 0 {
				answer[i] = stack.Top()
			}
		}

		stack.Push(nums2[i])
	}
	return answer
}

func main() {

}
