package main

// func findOriginalArray(changed []int) []int {
// 	sort.Slice(changed, func(i, j int) bool {
// 		return changed[i] < changed[j]
// 	})

// 	original := []int{}
// 	for len(changed) > 0 {
// 		dubled := false
// 		for j := 1; j < len(changed); j++ {
// 			if changed[0]*2 == changed[j] {
// 				original = append(original, changed[0])
// 				tempChanged := changed[j+1:]
// 				changed = changed[:j]
// 				changed = append(changed, tempChanged...)
// 				changed = changed[1:]
// 				dubled = true
// 				break
// 			}
// 		}
// 		if !dubled {
// 			break
// 		}
// 	}
// 	if len(changed) != 0 {
// 		return []int{}
// 	}
// 	return original
// }

// func main() {
// 	original := findOriginalArray([]int{1, 3, 4, 2, 6, 8})
// 	fmt.Println(original)
// }
