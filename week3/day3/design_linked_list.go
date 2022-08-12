package main

import "fmt"

type Node struct {
	Val  int
	next *Node
}
type MyLinkedList struct {
	Head *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	head := this.Head
	counter := 0
	for head != nil {
		if counter == index {
			return head.Val
		}
		counter++
		head = head.next
	}

	return -1

}

func (this *MyLinkedList) AddAtHead(val int) {
	newHead := Node{}
	newHead.Val = val
	newHead.next = this.Head
	this.Head = &newHead
}

func (this *MyLinkedList) AddAtTail(val int) {

	if this.Head == nil {
		this.Head = &Node{Val: val}
		return
	}

	head := this.Head
	for head.next != nil {
		head = head.next
	}
	head.next = &Node{Val: val}

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	newNode := Node{
		Val: val,
	}
	head := this.Head
	counter := 0
	if index == 0 {
		newNode.next = this.Head
		this.Head = &newNode
	}
	for head != nil {
		if counter == index-1 {
			newNode.next = head.next
			head.next = &newNode
			break
		}
		counter++
		head = head.next
	}

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		if this.Head.next == nil {
			this.Head = nil
		} else {
			this.Head = this.Head.next
		}

		return
	}
	head := this.Head
	counter := 0
	for head.next != nil {
		if counter == index-1 {
			head.next = head.next.next
			break
		}
		counter++
		head = head.next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
	obj := Constructor()
	obj.AddAtTail(1)
	// obj.AddAtHead(1)
	// obj.AddAtHead(3)
	// obj.AddAtIndex(1, 6)
	// obj.AddAtIndex(0, 4)
	// obj.AddAtTail(3)
	// obj.DeleteAtIndex(0)

	param_1 := obj.Get(0)
	fmt.Println(param_1)
}
