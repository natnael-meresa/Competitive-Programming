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

func AddAtTail(head *ListNode, val int) {

	if head == nil {
		head = &ListNode{Val: val}
		return
	}
	newHead := head
	for newHead.Next != nil {
		newHead = newHead.Next
	}
	newHead.Next = &ListNode{Val: val}
}

func AddList(head *ListNode, list *ListNode) {

	if head == nil {
		head = list
		return
	}
	newHead := head
	for newHead.Next != nil {
		newHead = newHead.Next
	}
	newHead.Next = list
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newHead := ListNode{}

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		newHead.Val = list1.Val
		list1 = list1.Next
	} else {
		newHead.Val = list2.Val
		list2 = list2.Next
	}

	for list1 != nil || list2 != nil {
		if list1 == nil {
			AddList(&newHead, list2)
			break
		} else if list2 == nil {
			AddList(&newHead, list1)
			break
		}

		if list1.Val <= list2.Val {
			AddAtTail(&newHead, list1.Val)
			list1 = list1.Next
		} else {
			AddAtTail(&newHead, list2.Val)
			list2 = list2.Next
		}
	}
	return &newHead

}

// func main() {

// 	node1 := ListNode{Val: 1, Next: nil}
// 	node2 := ListNode{Val: 2, Next: nil}
// 	node3 := ListNode{Val: 4, Next: nil}
// 	head := node1
// 	head.Next = &node2
// 	node2.Next = &node3

// 	list2node1 := ListNode{Val: 1, Next: nil}
// 	list2node2 := ListNode{Val: 3, Next: nil}
// 	list2node3 := ListNode{Val: 4, Next: nil}
// 	list2head := list2node1
// 	list2head.Next = &list2node2
// 	list2node2.Next = &list2node3

// 	newHead := mergeTwoLists(&head, &list2head)
// 	fmt.Println(newHead)
// 	fmt.Println(newHead.Next)
// 	fmt.Println(newHead.Next.Next)
// 	fmt.Println(newHead.Next.Next.Next)
// 	fmt.Println(newHead.Next.Next.Next.Next)

// }
