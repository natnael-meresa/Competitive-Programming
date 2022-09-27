package main

import "fmt"

func totalFruit(fruits []int) int {
	if len(fruits) < 2 {
		return 1
	}

	maxTotal := 0
	// lastIndex, fruit, total, firstIndex
	bascket2 := make([]int, 4)
	bascket1 := []int{0, fruits[0], 1, 0}
	j := 1
	for j < len(fruits) && fruits[j] == fruits[j-1] {
		bascket1[0]++
		bascket1[2]++
		j++
	}
	if j < len(fruits) {
		bascket2 = []int{j, fruits[j], 1, j}
	}
	maxTotal = bascket1[2] + bascket2[2]
	for i := j + 1; i < len(fruits); i++ {
		if fruits[i] == bascket1[1] {
			bascket1[2]++
			bascket1[0] = i
		} else if fruits[i] == bascket2[1] {
			bascket2[2]++
			bascket2[0] = i
		} else {
			if bascket1[0] > bascket2[0] {
				cI := bascket2[0] + 1
				lI := bascket1[3]

				for lI < cI {

					if fruits[lI] == bascket1[1] {
						bascket1[2] -= 1
					}
					lI++
				}
				bascket1[3] = cI
				bascket2 = []int{i, fruits[i], 1, i}
			} else {
				cI := bascket1[0] + 1
				lI := bascket2[3]
				for lI < cI {
					if fruits[lI] == bascket2[1] {
						bascket2[2] -= 1
					}
					lI++
				}
				bascket2[3] = cI
				bascket1 = []int{i, fruits[i], 1, i}
			}
		}

		if bascket1[2]+bascket2[2] > maxTotal {
			maxTotal = bascket1[2] + bascket2[2]
		}
	}

	return maxTotal
}

func main() {
	fmt.Println(totalFruit([]int{1, 0, 1, 4, 1, 4, 1, 2, 3}))
}
