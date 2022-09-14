package main

// import (
// 	"fmt"
// 	"math"
// )

// func myPow(x float64, n int) float64 {
// 	if x == 0 || x == 1 {
// 		return x
// 	}
// 	max := pow(x, int(math.Abs(float64(n))))
// 	if n < 0 {
// 		max = 1 / max
// 	}

// 	return max
// }

// func pow(x float64, n int) float64 {
// 	if n < 2 {
// 		return x
// 	}
// 	half := pow(x, n/2)
// 	res := half * half

// 	if n%2 != 0 {
// 		res *= x
// 	}
// 	return res
// }

// func main() {
// 	fmt.Println(myPow(2, 9))
// }
