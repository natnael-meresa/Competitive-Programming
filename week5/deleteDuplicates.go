package main

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

// func deleteDuplicates(head *ListNode) *ListNode {
// 	current := head
// 	var temp *ListNode

// 	for current != nil {
// 		dubl := false

// 		for current.Next != nil && current.Next.Val == current.Val {
// 			current.Next = current.Next.Next
// 			if !dubl {
// 				dubl = true
// 			}
// 		}

// 		if !dubl {
// 			if temp == nil {
// 				head = current
// 				temp = current
// 			} else {
// 				temp.Next = current
// 				temp = current
// 			}
// 		}
// 		current = current.Next
// 	}
// 	if temp == nil {
// 		return temp
// 	}
// 	temp.Next = nil
// 	return head
// }

// func main() {
// 	node1 := ListNode{Val: 1, Next: nil}
// 	node2 := ListNode{Val: 2, Next: nil}
// 	node3 := ListNode{Val: 3, Next: nil}
// 	node4 := ListNode{Val: 4, Next: nil}
// 	node5 := ListNode{Val: 4, Next: nil}
// 	node6 := ListNode{Val: 5, Next: nil}
// 	head := node1
// 	head.Next = &node2
// 	node2.Next = &node3
// 	node3.Next = &node4
// 	node4.Next = &node5
// 	node5.Next = &node6
// 	final := deleteDuplicates(&head)
// 	fmt.Println(final)
// 	fmt.Println(final.Next)
// 	fmt.Println(final.Next.Next)
// 	fmt.Println(final.Next.Next.Next)
// }

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	var temp *ListNode

	for current != nil {
		dubl := false
		for current.Next != nil && current.Next.Val == current.Val {
			if !dubl {
				dubl = true
			}
			current.Next = current.Next.Next
		}

		if dubl {
			if temp == nil {
				temp = current
			}
			temp.Next = current
		}

		current = current.Next

	}

	temp.Next = nil
	return temp
}
