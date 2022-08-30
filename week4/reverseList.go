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
	tempHead := head
	head1 := head

	for head1.Next != nil {
		head = head1.Next
		head1.Next = head.Next
		head.Next = tempHead
		tempHead = head
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
