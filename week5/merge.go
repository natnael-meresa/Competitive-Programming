package main

import "sort"

// import (
// 	"fmt"
// 	"sort"
// )

// func merge(intervals [][]int) [][]int {
// 	sort.Slice(intervals, func(i, j int) bool {
// 		return intervals[i][0] <= intervals[j][0]
// 	})
// 	merged := make([][]int, len(intervals))
// 	mergeNum := 0
// 	merged[0] = intervals[0]

// 	for i := 1; i < len(intervals); i++ {
// 		fmt.Println(merged)
// 		if intervals[i][0] <= merged[mergeNum][1] {
// 			if merged[mergeNum][1] <= intervals[i][1] {
// 				merged[mergeNum][1] = intervals[i][1]
// 			}
// 		} else {
// 			mergeNum++
// 			merged[mergeNum] = intervals[i]
// 		}
// 	}
// 	return merged
// }

// func main() {
// 	merged := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
// 	fmt.Println(merged)
// }

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	merged := [][]int{}
	merged = append(merged, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= merged[len(merged)-1][1] {
			if merged[len(merged)-1][1] <= intervals[i][1] {
				merged[len(merged)-1][1] = intervals[i][1]
			}
		} else {
			merged = append(merged, intervals[i])
		}
	}
	return merged
}
