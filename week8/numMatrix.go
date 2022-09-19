package main

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	for i := 0; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			matrix[i][j] = matrix[i][j-1] + matrix[i][j]
		}
	}
	return NumMatrix{
		matrix: matrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		if col1 == col2 {
			if col1 == 0 {
				sum += this.matrix[i][col1]
			} else {
				sum += this.matrix[i][col1] - this.matrix[i][col1-1]
			}
		} else if col1 != 0 {
			sum += this.matrix[i][col2] - this.matrix[i][col1-1]
		} else {
			sum += this.matrix[i][col2]
		}

	}

	return sum
}

// func main() {
// 	// matrix := [][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}}
// 	matrix := [][]int{{-4, -5}}
// 	obj := Constructor(matrix)

// 	// param_1 := obj.SumRegion(row1, col1, row2, col2)
// 	param_1 := obj.SumRegion(0, 1, 0, 1)

// 	fmt.Println(param_1)
// }

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
