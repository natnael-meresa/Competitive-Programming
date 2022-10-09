package main

import "fmt"

func carPooling(trips [][]int, capacity int) bool {
	prefix := make([]int, 1001)

	for i := 0; i < len(trips); i++ {
		numPassenger, from, to := trips[i][0], trips[i][1], trips[i][2]
		prefix[from] += numPassenger
		prefix[to] -= numPassenger
	}

	if prefix[0] > capacity {
		return false
	}
	for i := 1; i < 1001; i++ {

		prefix[i] += prefix[i-1]
		if prefix[i] > capacity {
			return false
		}
	}

	return true
}

func main() {
	trips := [][]int{{9, 0, 1}, {3, 3, 7}}
	fmt.Println(carPooling(trips, 4))
}
