package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.

 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// no global variable

func isPalindrome(head *ListNode) bool {
	_, ans := isPldrm(head, head)
	return ans
}

func isPldrm(left, right *ListNode) (*ListNode, bool) {
	if right == nil {
		return left, true
	}
	left, ans := isPldrm(left, right.Next)
	ans = ans && left.Val == right.Val
	return left.Next, ans
}

//
//
// global variable

// var left *ListNode

// func isPalindrome(head *ListNode) bool {
// 	left = head
// 	return isPldrm(head)
// }
// func isPldrm(right *ListNode) bool {
// 	if right == nil {
// 		return true
// 	}
// 	ans := isPldrm(right.Next) && left.Val == right.Val
// 	left = left.Next
// 	return ans
// }

// func isPalindrome(head *ListNode) bool {
// 	tempHead := head
// 	plStr := ""
// 	for tempHead != nil {
// 		plStr += strconv.Itoa(tempHead.Val)
// 		tempHead = tempHead.Next
// 	}

// 	for i, j := 0, len(plStr)-1; j > 0; j, i = j-1, i+1 {
// 		if plStr[i] != plStr[j] {
// 			return false
// 		}

// 	}

// 	return true
// }

func main() {
	node1 := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 1, Next: nil}
	node3 := ListNode{Val: 2, Next: nil}
	node4 := ListNode{Val: 1, Next: nil}
	node5 := ListNode{Val: 1, Next: nil}
	head := node1

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	fmt.Println(isPalindrome(&head))

}
