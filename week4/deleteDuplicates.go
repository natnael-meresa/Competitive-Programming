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

// -- this for unsorted values ---
// type uniqueValues map[int]struct{}

// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	u := uniqueValues{}
// 	temp := head
// 	u.add(temp.Val)
// 	for temp.Next != nil {
// 		if u.has(temp.Next.Val) {
// 			temp.Next = temp.Next.Next
// 		} else {
// 			u.add(temp.Next.Val)
// 			temp = temp.Next
// 		}
// 	}

// 	return head
// }

// func (u uniqueValues) has(val int) bool {
// 	_, ok := u[val]
// 	return ok
// }
// func (u uniqueValues) add(val int) {
// 	u[val] = struct{}{}
// }

func deleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil {
		for current.Next != nil && current.Next.Val == current.Val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}

	return head
}

// func main() {
// 	node1 := ListNode{Val: 2, Next: nil}
// 	node2 := ListNode{Val: 4, Next: nil}
// 	node3 := ListNode{Val: 1, Next: nil}
// 	node4 := ListNode{Val: 2, Next: nil}
// 	head := node1
// 	head.Next = &node2
// 	node2.Next = &node3
// 	node3.Next = &node4

// 	newNode := deleteDuplicates(&head)
// 	fmt.Println(newNode)
// 	fmt.Println(newNode.Next)
// 	fmt.Println(newNode.Next.Next)
// 	fmt.Println(newNode.Next.Next.Next)

// }
