package main

import (
	"strconv"
)

func calculate(s string) int {
	n := []int{}

	currNum := 0
	currentOp := "+"

	for _, x := range s {
		xs := string(x)
		xsNum, err := strconv.Atoi(xs)
		if err == nil {
			currNum = (currNum * 10) + xsNum
		} else if xs != " " {
			if currentOp == "+" {
				n = append(n, currNum)
			} else if currentOp == "-" {
				n = append(n, -currNum)
			} else if currentOp == "*" {
				tnum := n[len(n)-1]
				n = n[:len(n)-1]
				n = append(n, tnum*currNum)
			} else if currentOp == "/" {
				tnum := n[len(n)-1]
				n = n[:len(n)-1]
				n = append(n, tnum/currNum)
			}
			currentOp = xs
			currNum = 0
		}
	}

	if currentOp == "+" {
		n = append(n, currNum)
	} else if currentOp == "-" {
		n = append(n, -currNum)
	} else if currentOp == "*" {
		tnum := n[len(n)-1]
		n = n[:len(n)-1]
		n = append(n, tnum*currNum)
	} else if currentOp == "/" {
		tnum := n[len(n)-1]
		n = n[:len(n)-1]
		n = append(n, tnum/currNum)
	}

	result := 0

	for i := 0; i < len(n); i++ {
		result += n[i]
	}

	return result
}

// func main() {
// 	fmt.Println(calculate("3+2+2*1+1+4*2"))
// }
