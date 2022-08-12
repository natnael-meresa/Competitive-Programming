package main

func pivotIndex(nums []int) int {
	total := 0
	leftTotal := 0

	for _, x := range nums {
		total += x
	}

	for i, x := range nums {
		if leftTotal == (total - leftTotal - x) {
			return i
		}

		leftTotal += x
	}

	return -1
}
