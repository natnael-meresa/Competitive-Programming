package main

import "fmt"

func findTheWinner(n int, k int) int {
	friends := make([]int, n)

	for i := 0; i < n; i++ {
		friends[i] = i + 1
	}
	indexLoser := k
	for len(friends) > 1 {
		for indexLoser > len(friends) {
			indexLoser -= len(friends)
		}
		temp := friends[indexLoser:]
		friends = friends[:indexLoser-1]
		friends = append(friends, temp...)
		indexLoser += k - 1
	}

	return friends[0]
}

func main() {
	winner := findTheWinner(6, 5)
	fmt.Println(winner)
}
