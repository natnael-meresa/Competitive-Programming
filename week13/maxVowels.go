package main

import (
	"fmt"
	"strings"
)

// func maxVowels(s string, k int) int {
// 	maxVowel, curVowel, left := 0, 0, 0
// 	vowels := map[byte]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}}
// 	// vowels := "aeiou"

// 	for i := 0; i < k; i++ {
// 		if _, ok := vowels[s[i]]; ok {
// 			curVowel++
// 		}
// 	}

// 	for i := k; i < len(s); i++ {
// 		if _, ok := vowels[s[i]]; ok {
// 			curVowel++
// 		}

// 		if curVowel > maxVowel {
// 			maxVowel = curVowel
// 		}

// 		// if i-left == k-1 {
// 		if _, ok := vowels[s[left]]; ok {
// 			curVowel--
// 		}
// 		left++
// 		// }

// 	}

// 	return maxVowel
// }

func maxVowels(s string, k int) int {
	maxVowel, curVowel, left := 0, 0, 0
	// vowels := map[byte]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}}
	vowels := "aeiou"

	for i := 0; i < k-1; i++ {
		if ok := strings.Contains(vowels, string(s[i])); ok {
			curVowel++
		}
	}

	for i := k - 1; i < len(s); i++ {
		if ok := strings.Contains(vowels, string(s[i])); ok {
			curVowel++
		}

		if curVowel > maxVowel {
			maxVowel = curVowel
		}

		if ok := strings.Contains(vowels, string(s[left])); ok {
			curVowel--
		}
		left++
		// }

	}

	return maxVowel
}
func main() {
	fmt.Println(maxVowels("abciiidef", 3))
}
