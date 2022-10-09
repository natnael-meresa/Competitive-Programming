package main

// brute force way(naive way)
// func corpFlightBookings(bookings [][]int, n int) []int {
// 	ans := make([]int, n)

// 	for i := 0; i < len(bookings); i++ {
// 		for j := bookings[i][0]; j <= bookings[i][1]; j++ {
// 			ans[j-1] += bookings[i][2]
// 		}
// 	}

// 	return ans
// }

func corpFlightBookings(bookings [][]int, n int) []int {
	ans := make([]int, n)

	for _, booking := range bookings {
		ans[booking[0]-1] += booking[2]
		if booking[1] < n {
			ans[booking[1]] -= booking[2]
		}

	}

	for i := 1; i < n; i++ {
		ans[i] += ans[i-1]
	}

	return ans
}

// func main() {
// 	bookings := [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
// 	fmt.Println(corpFlightBookings(bookings, 5))
// }
