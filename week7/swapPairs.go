package main

// import "fmt"

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func swapPairs(head *ListNode) *ListNode {

// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	prvHead := swapPairs(head.Next.Next)
// 	tempHead := head
// 	head = head.Next

// 	head.Next = tempHead
// 	tempHead.Next = prvHead

// 	return head
// }

// func main() {

// 	node1 := ListNode{Val: 1, Next: nil}
// 	node2 := ListNode{Val: 2, Next: nil}
// 	node3 := ListNode{Val: 3, Next: nil}
// 	node4 := ListNode{Val: 4, Next: nil}
// 	node5 := ListNode{Val: 5, Next: nil}
// 	node6 := ListNode{Val: 6, Next: nil}
// 	head := node1

// 	head.Next = &node2
// 	node2.Next = &node3
// 	node3.Next = &node4
// 	node4.Next = &node5
// 	node5.Next = &node6
// 	newHead := swapPairs(&head)
// 	tempHead := newHead
// 	for tempHead != nil {
// 		fmt.Println(tempHead.Val)
// 		tempHead = tempHead.Next
// 	}
// }
