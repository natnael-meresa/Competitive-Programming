package main

import "fmt"

type MyQueue struct {
	items []int
}

func Constructor() MyQueue {
	myQueue := MyQueue{}
	return myQueue
}

func (this *MyQueue) Push(x int) {
	this.items = append(this.items, x)
}

func (this *MyQueue) Pop() int {
	last := this.items[0]
	this.items = this.items[1:]
	return last
}

func (this *MyQueue) Peek() int {
	return this.items[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.items) == 0
}

func main() {
	obj := Constructor()
	obj.Push(1)
	param_2 := obj.Pop()
	param_3 := obj.Peek()
	param_4 := obj.Empty()

	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(param_4)
}
