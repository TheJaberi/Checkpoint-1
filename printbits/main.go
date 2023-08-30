package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		PrintString("00000000")
		return
	}
	if num > 255 {
		PrintString("00000000")
		return
	}
	binNum := ToBin(num)
	PrintString(binNum)
}

func ToBin(num int) string {
	str := ""
	binMap := "01"
	for i := 0; i< 8; i++ {
		str = string(binMap[num%2]) + str
		num /= 2
	}
	return str
}

func PrintString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
