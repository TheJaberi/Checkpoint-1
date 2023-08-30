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
	args := os.Args[1] + os.Args[2]
	str := ""
	for _, r := range args {
		for _, s := range str {
			if s == r {
				goto there
			}
		}
		str += string(r)
	there:
	}
	PrintString(str)
}

func PrintString(str string){
	for _, r := range str {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
