package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args[1:]) != 2 {
		z01.PrintRune('\n')
		return
	}
	str := ""
	for _, r := range os.Args[1] {
		if containsRune(os.Args[2],r) && !containsRune(str,r) {
			str += string(r)
		}
	}
	PrintString(str)
}

func PrintString(str string) {
	for _, r := range str {
		z01.PrintRune(r)
	}
}

func containsRune(str string, char rune) bool {
	for _, c := range str {
		if c == char {
			return true
		}
	}
	return false
}
