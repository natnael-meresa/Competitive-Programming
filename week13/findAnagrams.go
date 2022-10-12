package main

import "fmt"

func findAnagrams(s string, p string) []int {
	ans := []int{}

	anagramHash := make(map[byte]int, len(p))
	for i := 0; i < len(p); i++ {
		anagramHash[p[i]]++
	}

	left, count := 0, len(p)
	for i := 0; i < len(s); i++ {

		if _, ok := anagramHash[s[i]]; ok {
			anagramHash[s[i]]--
			if anagramHash[s[i]] >= 0 {
				count--
			}
		}

		if count == 0 {
			ans = append(ans, left)
		}
		fmt.Println(anagramHash, "i", i, count)
		if i-left == len(p)-1 {
			if _, ok := anagramHash[s[left]]; ok {
				anagramHash[s[left]]++
				if anagramHash[s[left]] > 0 {
					count++
				}
			}
			left++
		}
	}

	return ans
}

func main() {
	fmt.Println(findAnagrams("baa", "aa"))
}
