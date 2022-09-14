package main

import (
	"fmt"
	"strings"
)

// posible solution but it can not pass b/c of time limit
func kthLargestNumber(nums []string, k int) string {
	for i := 0; i < k; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if len(nums[j]) > len(nums[j+1]) {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			} else if len(nums[j]) == len(nums[j+1]) {
				jL := strings.Split(nums[j], "")
				jPL := strings.Split(nums[j+1], "")
				for s := 0; s < len(nums[j]); s++ {
					if jL[s] != jPL[s] {
						if jL[s] > jPL[s] {
							nums[j], nums[j+1] = nums[j+1], nums[j]
						}
						break
					}
				}
			}
		}
	}
	return nums[len(nums)-k]
}

func main() {

	// kth := kthLargestNumber([]string{"683339452288515879", "7846081062003424420", "4805719838", "4840666580043", "83598933472122816064", "522940572025909479", "615832818268861533", "65439878015", "499305616484085", "97704358112880133", "23861207501102", "919346676", "60618091901581", "5914766072", "426842450882100996", "914353682223943129", "97", "241413975523149135", "8594929955620533", "55257775478129", "528", "5110809", "7930848872563942788", "758", "4", "38272299275037314530", "9567700", "28449892665", "2846386557790827231", "53222591365177739", "703029", "3280920242869904137", "87236929298425799136", "3103886291279"}, 3)
	kth := kthLargestNumber([]string{"423", "521", "2", "42"}, 2)
	str := "423"
	fmt.Println("wat", str[0])
	fmt.Println()
	fmt.Println(kth)
}

// func kthLargestNumber(nums []string, k int) string {
//     sort.Slice(nums, func(i,j int) bool {
//         if len(nums[i]) == len(nums[j]){
//             return nums[i] < nums[j]
//         } else {
//             return len(nums[i]) < len(nums[j])
//         }
//     })
//     return nums[len(nums)-k]
// }
