package main

import "math"

func isPowerOfFour(n int) bool {
	if n == 0 {
		return false
	}

	float := math.Log(float64(n)) / math.Log(4)
	return float32(float) == float32(math.Trunc(float))

}

func main() {
	isIt := isPowerOfFour(7)
	print(isIt)
}
