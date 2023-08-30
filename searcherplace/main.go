package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		z01.PrintRune('\n')
		return
	}
	str := []rune(os.Args[1])
	from := os.Args[2]
	to := os.Args[3]
	if len(from) > 1 || len(to) > 1 {
		z01.PrintRune('\n')
		return
	}
	for i, r := range str {
		if r == rune(from[0]) {
			str[i] = rune(to[0])
		}
	}
	PrintStringln(string(str))
}

func PrintStringln(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
