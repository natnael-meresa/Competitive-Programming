package main

import "fmt"

type MinStack struct {
	items []int
}

func Constructor() MinStack {
	item := MinStack{}
	return item
}

func (this *MinStack) Push(val int) {
	this.items = append(this.items, val)
}

func (this *MinStack) Pop() {
	this.items = this.items[:len(this.items)-1]
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	min := this.items[0]

	for _, x := range this.items {
		if min > x {
			min = x
		}
	}

	return min
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(-2)
	obj.Push(0)
	obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.GetMin()

	fmt.Println(param_3)
	fmt.Println(param_4)
}
