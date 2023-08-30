package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args[1:]) != 2 {
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]
	if IsHidden(str1,str2) {
		z01.PrintRune('1')
	} else {
		z01.PrintRune('0')
	}
}

// func IsHidden(s1, s2 string) bool {
// 	if s1 == "" {
// 		return true
// 	}
// 	start := 0
// 	count := 0
// 	for _, s := range s1 {
// 		for i := start + 1; i < len(s2); i++ {
// 			if s == rune(s2[i]) {
// 				start = i
// 				count++
// 				break
// 			}
// 		}
// 	}
// 	if (s2[start] == s1[len(s1)-1] && count == len(s1)) {
// 		return true
// 	}
// 	return false
// }

// or 
func IsHidden(s1,s2 string) bool {
	if s1 == "" {
		return true
	}
	s1index := 0
	for _, r := range s2 {
		if r == rune(s1[s1index]) {
			s1index++
			if s1index == len(s1) {
				return true
			}
		}
	}
	return false
}