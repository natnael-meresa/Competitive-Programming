package main

import (
	"fmt"
	"sort"
)

func carFleet(target int, position []int, speed []int) int {
	pos_speed := make(map[int]int)
	mst := make([]int, 0)

	for x := range position {
		pos_speed[position[x]] = speed[x]
	}
	sort.Ints(position)

	for _, x := range position {
		for len(mst) != 0 && numItration(target, x, pos_speed[x]) >= numItration(target, mst[len(mst)-1], pos_speed[mst[len(mst)-1]]) {
			mst = mst[:len(mst)-1]
		}

		mst = append(mst, x)
	}

	return len(mst)

}

func numItration(t, p, s int) float32 {
	return (float32(t) - float32(p)) / float32(p)
}

func main() {
	flet := carFleet(12, []int{10, 8, 0}, []int{3, 2})
	fmt.Println(flet)
}

// slow
// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func carFleet(target int, position []int, speed []int) int {
// 	pos_speed_itr := make(map[int][]float32)
// 	mst := make([]int, 0)

// 	for x := range position {
// 		pos_speed_itr[position[x]] = []float32{float32(speed[x]), (float32(target) - float32(position[x])) / float32(speed[x])}
// 	}
// 	sort.Ints(position)

// 	for _, x := range position {
// 		for len(mst) != 0 && pos_speed_itr[x][1] >= pos_speed_itr[mst[len(mst)-1]][1] {
// 			mst = mst[:len(mst)-1]
// 		}

// 		mst = append(mst, x)
// 	}

// 	return len(mst)

// }

// func main() {
// 	flet := carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3})

// 	fmt.Println(flet)
// }
