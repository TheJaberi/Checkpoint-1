package main

import (
	"fmt"
	"test"
)

func main() {
	input := byte(0b00100110) // 0010 0110
	output := piscine.ReverseBits(input)
	fmt.Printf("%08b\n", input)  // 0010 0110
	fmt.Printf("%08b\n", output) // 0110 0100
}
