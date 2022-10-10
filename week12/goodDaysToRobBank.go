package main

func goodDaysToRobBank(security []int, time int) []int {
	res := []int{}
	n := len(security)
	if time == 0 {
		for i := 0; i < n; i++ {
			res = append(res, i)
		}
		return res
	}

	count := 0
	left := make([]int, n)
	right := make([]int, n)

	for i := 1; i < n; i++ {
		if security[i] <= security[i-1] {
			count++
			left[i] = count
		} else {
			count = 0
		}
	}
	count = 0
	for i := n - 2; i >= 0; i-- {
		if security[i] <= security[i+1] {
			count++
			right[i] = count
		} else {
			count = 0
		}
	}

	for i := 1; i < n-1; i++ {
		if left[i] >= time && right[i] >= time {
			res = append(res, i)
		}
	}

	return res
}

func main() {

}
