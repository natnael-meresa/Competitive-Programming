package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {

	node.Val = node.Next.Val
	newHead := node
	for newHead.Next.Next != nil {
		newHead = newHead.Next
		newHead.Val = newHead.Next.Val
	}
	newHead.Next = nil
}

func main() {
	ls := ListNode{}
	deleteNode(&ls)
}
