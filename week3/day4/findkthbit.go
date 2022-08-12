package main

import (
	"fmt"
)

func findKthBit(n int, k int) byte {
	final := s(n)
	fmt.Println(string(final))

	return final[k-1]
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)

}

func invert(s string) string {
	r := []rune(s)
	for i := range s {
		if r[i] == '0' {
			r[i] = '1'
		} else {
			r[i] = '0'
		}
	}
	return string(r)
}

func s(n int) string {
	if n == 1 {
		return "0"
	}
	return s(n-1) + "1" + reverse(invert(s(n-1)))
}

func main() {
	fmt.Println(string(findKthBit(6, 11)))
}
