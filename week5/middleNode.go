package main

// import "fmt"

// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	middleHead, endHead, endPostion := head, head, 0

	for endHead != nil {
		endPostion++

		if endPostion%2 == 0 {
			middleHead = middleHead.Next
		}
		endHead = endHead.Next
	}
	return middleHead
}

// func main() {
// 	node1 := ListNode{Val: 1, Next: nil}
// 	node2 := ListNode{Val: 2, Next: nil}
// 	node3 := ListNode{Val: 3, Next: nil}
// 	node4 := ListNode{Val: 4, Next: nil}
// 	node5 := ListNode{Val: 5, Next: nil}
// 	head := node1
// 	head.Next = &node2
// 	node2.Next = &node3
// 	node3.Next = &node4
// 	node4.Next = &node5

// 	middle := middleNode(&head)
// 	fmt.Println(middle)
// 	fmt.Println(middle.Next)
// }
