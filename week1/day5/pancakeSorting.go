package main

import "fmt"

func pancakeSort(arr []int) []int {
	counts, n, tobe := []int{}, len(arr), []int{}
	for i := 0; i < n; i++ {
		max := 0
		for j := 0; j < n-i; j++ {
			if arr[max] < arr[j] {
				max = j
			}
		}

		if max == n-(i+1) {
			continue
		}
		if max > 0 {
			tobe = arr[:max+1]
			for i, j := 0, len(tobe)-1; i < j; i, j = i+1, j-1 {
				tobe[i], tobe[j] = tobe[j], tobe[i]
			}
			counts = append(counts, max+1)
		}

		if arr[0] > arr[1] {
			counts = append(counts, n-i)
			tobe = arr[:(n - i)]
			for i, j := 0, len(tobe)-1; i < j; i, j = i+1, j-1 {
				tobe[i], tobe[j] = tobe[j], tobe[i]
			}
		}

	}

	return counts
}

func main() {

	final := pancakeSort([]int{3, 2, 4, 1})

	fmt.Println(final)
}
