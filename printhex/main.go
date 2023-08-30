package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	hexstring := ToHex(number)
	PrintStringln(hexstring)
	// or 
	// hexstring += "\n"
	// os.Stdout.WriteString(hexstring)
}

func ToHex(n int) string {
	hexmap := "0123456789abcdef"
	str := ""
	for n != 0 {
		str = string(hexmap[n%16]) + str
		n /= 16
	}
	return str
}

func PrintStringln(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

// or
// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}
// 	number, err := strconv.Atoi(os.Args[1])
// 	if err != nil {
// 		fmt.Println("ERROR")
// 		return
// 	}
// 	fmt.Println(fmt.Sprintf("%x",number))
// } 