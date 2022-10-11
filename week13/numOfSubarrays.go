package main

func numOfSubarrays(arr []int, k int, threshold int) int {
	count := 0
	left := 0
	for i := 1; i < k; i++ {
		arr[i] += arr[i-1]
	}

	if arr[k-1]/k >= threshold {
		count++
	}

	for i := k; i < len(arr); i++ {
		arr[i] += arr[i-1]

		if (arr[i]-arr[left])/k >= threshold {
			count++
		}
		left++
	}

	return count
}
