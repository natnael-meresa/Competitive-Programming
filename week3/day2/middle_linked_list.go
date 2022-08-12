package main

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	total := 1
	totalHead := head
	for totalHead.Next != nil {
		totalHead = totalHead.Next
		total++
	}

	for i := 0; i < int(math.Ceil(float64(total/2))); i++ {
		head = head.Next
	}

	return head

}
