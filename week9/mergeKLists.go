package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n < 1 {
		return nil
	}
	j := 0
	for j < len(lists) && lists[j] == nil {
		j++
	}
	if j >= len(lists) {
		return nil
	}
	list := lists[j]

	for i := 1; i < n; i++ {
		if lists[i] != nil {
			list = merge2Lists(list, lists[i])
		}
	}

	return list
}

func merge2Lists(l1, l2 *ListNode) *ListNode {
	h := l1
	if l2.Val <= l1.Val {
		l1 = l2
		tempL2 := l2.Next
		l1.Next = h
		h = l1
		l2 = tempL2
	}

	for h.Next != nil && l2 != nil {

		if l2.Val <= h.Next.Val {
			tempH := h.Next
			tmpL2 := l2.Next
			h.Next = l2
			h.Next.Next = tempH
			l2 = tmpL2
		}
		h = h.Next
	}

	if l2 != nil {
		h.Next = l2
	}

	return l1
}

// func main() {

// 	// node1 := ListNode{Val: 1, Next: nil}
// 	// node2 := ListNode{Val: 4, Next: nil}
// 	// node3 := ListNode{Val: 5, Next: nil}
// 	// head := node1
// 	// head.Next = &node2
// 	// node2.Next = &node3

// 	// h2node1 := ListNode{Val: 1, Next: nil}
// 	// h2node2 := ListNode{Val: 3, Next: nil}
// 	// h2node3 := ListNode{Val: 4, Next: nil}
// 	// h2head := h2node1
// 	// h2head.Next = &h2node2
// 	// h2node2.Next = &h2node3

// 	// h3node1 := ListNode{Val: 2, Next: nil}
// 	// h3node2 := ListNode{Val: 6, Next: nil}
// 	// h3head := h3node1
// 	// h3head.Next = &h3node2

// 	// newNode := mergeKLists([]*ListNode{&head, &h2head})
// 	// fmt.Println(newNode)
// 	// fmt.Println(newNode.Next)
// 	// fmt.Println(newNode.Next.Next)
// 	// fmt.Println(newNode.Next.Next.Next)
// 	// fmt.Println(newNode.Next.Next.Next.Next)
// 	// fmt.Println(newNode.Next.Next.Next.Next.Next)
// 	// fmt.Println(newNode.Next.Next.Next.Next.Next.Next)

// 	node1 := ListNode{Val: 1, Next: nil}

// 	head2 := node1

// 	newNode := mergeKLists([]*ListNode{nil, &head2})
// 	fmt.Println(newNode)

// }
