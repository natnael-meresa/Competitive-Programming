package main

func xorQueries(arr []int, queries [][]int) []int {
	ans := make([]int, len(queries))
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[i-1] ^ arr[i]
	}

	for i, query := range queries {
		left := query[0]
		right := query[1]
		val := 0
		if left != 0 {
			val = arr[right] ^ arr[left-1]
		} else {
			val = arr[right]
		}

		ans[i] = val
	}

	return ans
}

// func main() {
// 	arr := []int{1, 3, 4, 8}
// 	queries := [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}

// 	fmt.Println(xorQueries(arr, queries))
// }
