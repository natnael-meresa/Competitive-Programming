package main

func getDescentPeriods(prices []int) int64 {
	if len(prices) == 1 {
		return 1
	}
	left := 0
	sum := 0
	i := 1
	for ; i < len(prices); i++ {
		if prices[i-1]-prices[i] != 1 {
			n := i - left
			sum += (n * (n + 1)) / 2
			left = i
		}
	}

	n := i - left
	sum += (n * (n + 1)) / 2
	left = i

	return int64(sum)
}

// func main() {
// 	fmt.Println(getDescentPeriods([]int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3}))
// }
