package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	endstr := ""
	str := os.Args[1]
	vowelfound := false
	for i, r := range str {
		if IsVowel(r) {
			str = str[i:] + endstr + "ay"
			vowelfound = true
			break
		} else {
			endstr += string(r)
		}
	}
	if !vowelfound {
		PrintStringln("No vowels")
		return
	}
	PrintStringln(str)
}

func PrintStringln(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func IsVowel(r rune) bool {
	vowels := "aeiouAEIOU"
	for _, v := range vowels {
		if v == r {
			return true
		}
	}
	return false
}
