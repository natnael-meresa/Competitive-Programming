package main

type ValidStack struct {
	items []int
}

func Constructor() ValidStack {
	item := ValidStack{}
	return item
}

func (this *ValidStack) Push(val int) {
	this.items = append(this.items, val)
}

func (this *ValidStack) Pop() {
	this.items = this.items[:len(this.items)-1]
}

func (this *ValidStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *ValidStack) Empty() bool {
	return len(this.items) == 0
}

func validateStackSequences(pushed []int, popped []int) bool {
	pushPointer := 0
	popPointer := 0
	stack := Constructor()
	stack.Push(pushed[pushPointer])
	pushPointer++
	for pushPointer <= len(pushed) && popPointer < len(popped) {
		if !stack.Empty() && stack.Top() == popped[popPointer] {
			stack.Pop()
			popPointer++
		} else {

			if pushPointer >= len(pushed) {
				break
			}
			stack.Push(pushed[pushPointer])
			pushPointer++

		}

	}

	return stack.Empty()
}

func main() {

	validateStackSequences([]int{22, 3, 3}, []int{1, 23, 3})

}
