package main

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

func dailyTemperatures(temperatures []int) []int {

	stack := Constructor()
	ans := make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {

		for len(stack.items) != 0 && temperatures[stack.Top()] <= temperatures[i] {
			stack.Pop()
		}

		if len(stack.items) == 0 {
			ans[i] = 0
		} else {
			ans[i] = stack.Top() - i
		}
		stack.Push(i)
	}

	return ans

}

func main() {
	dailyTemperatures([]int{1, 2, 3})
}
