package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args[1:]) != 1 {
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	for n%2 == 0 {
		n /= 2
	}
	if n == 1 {
		PrintString("true")
	} else {
		PrintString("false")
	}
}

func PrintString(str string) {
	for _, r := range str {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
