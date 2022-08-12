package main

import "strconv"

func compress(chars []byte) int {
	num, firstPointer := 0, 0
	for secondPointer := 0; secondPointer <= len(chars); secondPointer++ {
		if secondPointer == len(chars) || chars[firstPointer] != chars[secondPointer] {
			chars[num] = chars[firstPointer]
			num++
			length := secondPointer - firstPointer
			if length > 1 {
				for _, digit := range strconv.Itoa(length) {
					chars[num] = byte(digit)
					num++
				}
			}
			firstPointer = secondPointer
		}
	}
	return num
}

// func main() {
// 	fmt.Println(compress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}))
// 	// fmt.Println(compress([]byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'}))
// }
