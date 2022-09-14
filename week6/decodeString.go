package main

// import (
// 	"fmt"
// 	"strconv"
// 	"unicode"
// )

// // func decodeString(s string) string {
// // 	stackTime := []int{}
// // 	stackStrings := []string{}
// // 	str := strings.Split(s, "")
// // 	finalStr := ""
// // 	tempStr := ""
// // 	lcStr := ""
// // 	lcNum := ""
// // 	for i := 0; i < len(str); i++ {

// // 		switch str[i] {
// // 		case "[":
// // 			newLcNum, _ := strconv.Atoi(lcNum)
// // 			stackTime = append(stackTime, newLcNum)
// // 			lcNum = ""
// // 			if lcStr != "" {
// // 				if string(lcStr[len(lcStr)-1]) != str[i-1] {
// // 					tempStr += lcStr
// // 					fmt.Println("what")
// // 				} else {
// // 					stackStrings = append(stackStrings, lcStr)
// // 				}
// // 				lcStr = ""
// // 			}

// // 		case "]":
// // 			if lcStr != "" {
// // 				fmt.Println(stackTime)

// // 				stackStrings = append(stackStrings, lcStr)
// // 				lcStr = ""
// // 			}

// // 			if len(stackTime) == 1 {
// // 				tmpFinal := ""
// // 				if tempStr != "" {
// // 					tmpFinal = stackStrings[0] + tempStr
// // 				}

// // 				for i := 0; i < stackTime[0]; i++ {
// // 					if tmpFinal != "" {
// // 						finalStr += tmpFinal
// // 					} else {
// // 						finalStr += stackStrings[0]
// // 					}
// // 				}

// // 				stackTime = stackTime[:len(stackTime)-1]
// // 				stackStrings = []string{}

// // 			} else if len(stackTime) > 1 {
// // 				for i := 0; i < stackTime[len(stackTime)-1]; i++ {
// // 					if len(stackStrings) > 1 {

// // 						tempStr += stackStrings[len(stackStrings)-1]
// // 					}
// // 				}

// // 				stackTime = stackTime[:len(stackTime)-1]
// // 			}
// // 		default:
// // 			_, err := strconv.Atoi(str[i])
// // 			if err == nil {
// // 				lcNum += str[i]
// // 				if len(stackTime) < 1 && lcStr != "" {
// // 					finalStr += lcStr
// // 					lcStr = ""
// // 				}
// // 			} else {
// // 				lcStr += str[i]
// // 			}
// // 		}
// // 	}

// // 	fmt.Println(stackTime)
// // 	if lcStr != "" {
// // 		finalStr += lcStr
// // 	}
// // 	return finalStr
// // }

// type StrStack []string

// func decodeString(s string) string {
// 	strStack := []string{}
// 	final := ""

// 	for _, x := range s {
// 		char := string(x)

// 		if char != "]" {
// 			strStack = append(strStack, char)
// 		} else {
// 			item := ""

// 			for len(strStack) > 0 && strStack[len(strStack)-1] != "[" {

// 				tempItem := strStack[len(strStack)-1]
// 				strStack = strStack[:len(strStack)-1]
// 				item = tempItem + item
// 			}

// 			if len(strStack) > 0 {
// 				strStack = strStack[:len(strStack)-1]
// 			}
// 			num := ""
// 			for len(strStack) > 0 && isNum(strStack[len(strStack)-1]) {
// 				tempItem := strStack[len(strStack)-1]
// 				strStack = strStack[:len(strStack)-1]
// 				num = tempItem + num
// 			}
// 			k, _ := strconv.ParseInt(num, 10, 32)
// 			subStr := ""

// 			for i := 0; i < int(k); i++ {
// 				subStr += item
// 			}

// 			strStack = append(strStack, subStr)

// 		}

// 	}

// 	for len(strStack) > 0 {
// 		tempFinal := strStack[len(strStack)-1]
// 		strStack = strStack[:len(strStack)-1]
// 		final = tempFinal + final
// 	}
// 	return final
// }

// func isNum(s string) bool {
// 	for _, c := range s {
// 		if !unicode.IsDigit(c) {
// 			return false
// 		}
// 	}
// 	return true
// }

// func isDigit(s string) bool {
// 	for _, c := range s {
// 		if !unicode.IsDigit(c) {
// 			return false
// 		}
// 	}
// 	return true
// }
// func main() {
// 	fmt.Println(decodeString("3[z]2[2[y]pq4[2[jk]e1[f]]]ef"))
// }
