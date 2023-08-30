package main

import (
	"fmt"
	"unicode"
)

func PrintMemory(arr [10]byte) {
	for i, b := range arr {
		fmt.Printf("%02x ", b)
		if (i+1)%4 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()

	for _, b := range arr {
		r := rune(b)
		if unicode.IsGraphic(r) {
			fmt.Printf("%c", r)
		} else {
			fmt.Printf(".")
		}
	}
	fmt.Println()
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
