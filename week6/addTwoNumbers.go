package main

import (
	"fmt"
	"math/big"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	left, right := "", ""
	tempHead1, tempHead2 := l1, l2
	for tempHead1 != nil {
		left = strconv.Itoa(tempHead1.Val) + left
		tempHead1 = tempHead1.Next
	}

	for tempHead2 != nil {
		right = strconv.Itoa(tempHead2.Val) + right
		tempHead2 = tempHead2.Next
	}

	leftNum, _ := new(big.Int).SetString(left, 0)
	rightNum, _ := new(big.Int).SetString(right, 0)
	final := (leftNum.Add(leftNum, rightNum)).String()
	cur, _ := strconv.Atoi(string(final[len(final)-1]))
	finalL := &ListNode{
		Val: cur,
	}
	temp := finalL
	for i := len(final) - 1; i >= 0; i-- {
		cur, _ := strconv.Atoi(string(final[i]))
		curL := ListNode{
			Val: cur,
		}
		temp.Next = &curL
		temp = temp.Next
	}
	return finalL
}
func main() {
	node1 := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: nil}
	node3 := ListNode{Val: 3, Next: nil}
	node4 := ListNode{Val: 4, Next: nil}
	node5 := ListNode{Val: 5, Next: nil}
	head := node1

	h2Node1 := ListNode{Val: 1, Next: nil}
	h2Node2 := ListNode{Val: 2, Next: nil}
	h2Node3 := ListNode{Val: 3, Next: nil}
	h2Node4 := ListNode{Val: 4, Next: nil}
	h2Node5 := ListNode{Val: 5, Next: nil}
	h2Head := h2Node1

	head.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	h2Head.Next = &h2Node2
	h2Node2.Next = &h2Node3
	h2Node3.Next = &h2Node4
	h2Node4.Next = &h2Node5

	newHead := addTwoNumbers(&head, &h2Head)
	fmt.Println(newHead)
	fmt.Println(newHead.Next)
	fmt.Println(newHead.Next.Next)
}

// [1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
// [5,6,4]

// 9223372036854775807
