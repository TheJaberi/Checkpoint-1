package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args[1:]) != 1 {
		z01.PrintRune('\n')
		return
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		z01.PrintRune('\n')
		return
	}
	for i := 1; i <= 9; i++ {
		z01.PrintRune(rune(i + '0'))
		z01.PrintRune(' ')
		z01.PrintRune('x')
		z01.PrintRune(' ')
		PrintString(os.Args[1])
		z01.PrintRune(' ')
		z01.PrintRune('=')
		z01.PrintRune(' ')
		PrintNum(i * num)
		z01.PrintRune('\n')
	}
}

func PrintNum(n int) {
	str := ""
	for n != 0 {
		str = string(n % 10 + '0') + str
		n /= 10
	}
	PrintString(str)
}

func PrintString(str string) {
	for _, r := range str {
		z01.PrintRune(r)
	}
}
