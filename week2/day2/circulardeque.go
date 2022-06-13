package main

import "fmt"

type MyCircularDeque struct {
	items []int
	cap   int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		items: make([]int, 0, k),
		cap:   k,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	this.items = append([]int{value}, this.items...)
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	this.items = append(this.items, value)

	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	this.items = append(this.items[:0], this.items[0+1:]...)

	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	this.items = this.items[:len(this.items)-1]

	return true

}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}

	return this.items[0]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.items[len(this.items)-1]
}

func (this *MyCircularDeque) IsEmpty() bool {
	if len(this.items) == 0 {
		return true
	}

	return false
}

func (this *MyCircularDeque) IsFull() bool {
	if len(this.items) == this.cap {
		return true
	}

	return false
}

func main() {
	k := 8
	obj := Constructor(k)
	param_1 := obj.InsertFront(5)
	fmt.Println(param_1)
	param_5 := obj.GetFront()
	fmt.Println(param_5)

	param_7 := obj.IsEmpty()
	fmt.Println(param_7)

	param_3 := obj.DeleteFront()
	fmt.Println(param_3)

	param_2 := obj.InsertLast(3)
	fmt.Println(param_2)

	// param_4 := obj.DeleteLast()
	// param_5 := obj.GetFront()
	// param_6 := obj.GetRear()
	// param_7 := obj.IsEmpty()
	// param_8 := obj.IsFull()

}
