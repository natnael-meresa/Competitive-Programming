package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// func pairSum(head *ListNode) int {
// 	arr := []int{}
// 	tmpH := head
// 	maxTwin, currentSum := 0, 0
// 	for tmpH != nil {
// 		arr = append(arr, tmpH.Val)
// 		tmpH = tmpH.Next
// 	}
// 	n := len(arr)
// 	for i := 0; i <= (n/2 - 1); i++ {
// 		currentSum = arr[i] + arr[n-1-i]
// 		if currentSum > maxTwin {
// 			maxTwin = currentSum
// 		}
// 	}

// 	return maxTwin
// }

func pairSum(head *ListNode) int {
	fast := head
	slow := head
	fMoves := 0

	for fast != nil {
		fast = fast.Next
		fMoves++
		if fMoves%2 == 0 {
			slow = slow.Next
		}
	}

	prv := slow
	nxt := slow

	for nxt.Next != nil {
		slow = nxt.Next
		nxt.Next = slow.Next
		slow.Next = prv
		prv = slow
	}

	max := 0
	for slow != nil {
		if head.Val+slow.Val > max {
			max = head.Val + slow.Val
		}

		head = head.Next
		slow = slow.Next
	}

	return max

}

func main() {

}
