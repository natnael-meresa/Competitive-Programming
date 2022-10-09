package main

import "fmt"

// slow
// func chalkReplacer(chalk []int, k int) int {

//     i :=0;
//     for k > 0 {
//         if k < chalk[i]{
//             return i
//         }
//         k -= chalk[i]
//         i++
//         if i >= len(chalk){
//             i = 0
//         }
//     }
//     return i
// }

// in this approach first find the sum of all the element of the initial array to compare it with the value of the k
// then find after some rotations at which index does the pointer stops in which the student need to change the chalk .
// then simply return that pointer
func chalkReplacer(chalk []int, k int) int {
	sum := 0
	for i := 0; i < len(chalk); i++ {
		sum += chalk[i]
	}

	k = k % sum
	i := 0
	fmt.Println(k)
	for chalk[i] <= k {
		k -= chalk[i]
		i = (i + 1) % len(chalk)
	}

	return i
}

// func main() {
// 	chalk := []int{5, 1, 5}
// 	fmt.Println(chalkReplacer(chalk, 10))
// }
