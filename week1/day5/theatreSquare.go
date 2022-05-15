package main

import "fmt"

func squre(n, m, a int) int64 {

	row := n / a
	column := m / a

	if n%a > 0 {
		row++
	}

	if m%a > 0 {
		column++
	}

	return int64(int64(row) * int64(column))
}
func main() {
	var n int
	var m int
	var a int
	fmt.Scan(&n, &m, &a)
	result := squre(n, m, a)

	fmt.Println(result)

}
