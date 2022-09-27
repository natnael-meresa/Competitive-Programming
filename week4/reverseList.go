package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return head
	}
	prv := head
	nxt := head

	for nxt.Next != nil {
		head = nxt.Next
		nxt.Next = head.Next
		head.Next = prv
		prv = head
	}

	return head

}

// func main() {
// 	node1 := ListNode{Val: 1, Next: nil}
// 	node2 := ListNode{Val: 2, Next: nil}
// 	node3 := ListNode{Val: 3, Next: nil}
// 	head := node1
// 	head.Next = &node2
// 	node2.Next = &node3
// 	fmt.Println(head)

// 	newHead := reverseList(&head)
// 	fmt.Println(newHead)
// 	fmt.Println(newHead.Next)
// 	fmt.Println(newHead.Next.Next)
// }
