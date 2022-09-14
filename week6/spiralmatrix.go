package main

// import "fmt"

// func spiralOrder(matrix [][]int) []int {
// 	newMatrix := []int{}
// 	fR, fC, lR, lC := 0, 0, len(matrix)-1, len(matrix[0])-1
// 	if lR == 0 {
// 		for i := 0; i <= lC; i++ {
// 			newMatrix = append(newMatrix, matrix[0][i])
// 		}
// 		return newMatrix
// 	}
// 	if lC == 0 {
// 		for i := 0; i <= lR; i++ {
// 			newMatrix = append(newMatrix, matrix[i][0])
// 		}
// 		return newMatrix
// 	}
// 	for lR >= fR && lC >= fC {

// 		for i := fC; i <= lC; i++ {
// 			newMatrix = append(newMatrix, matrix[fR][i])
// 		}

// 		fR++

// 		if lR >= fR {
// 			fmt.Println(fC, lC)
// 			for i := fR; i <= lR; i++ {
// 				newMatrix = append(newMatrix, matrix[i][lC])
// 			}

// 			lC--
// 		}

// 		if lR >= fR {
// 			for i := lC; i >= fC; i-- {
// 				newMatrix = append(newMatrix, matrix[lR][i])
// 			}
// 			lR--
// 		}

// 		if lR > fR {
// 			for i := lR; i >= fR; i-- {
// 				newMatrix = append(newMatrix, matrix[i][fC])
// 			}
// 			fC++

// 		}
// 	}

// 	return newMatrix
// }
// func main() {
// 	fmt.Println(spiralOrder([][]int{{1, 11}, {2, 12}, {3, 13}, {4, 14}, {5, 15}, {6, 16}, {7, 17}, {8, 18}, {9, 19}, {10, 2}}))
// }

// 0, ...
// ...., lastcolumn
// lastrow, ...
// ...., 0

// 1, ...
// ... , lastcolumn-1
// lastrow - 1 , ...
// ...., 1
