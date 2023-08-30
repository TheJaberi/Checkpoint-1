package main

import (
	"fmt"
	"os"
)

func main() {
	options := 0

	for _, arg := range os.Args[1:] {
		if arg == "-h" {
			fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
			return
		} else if len(arg) > 0 && arg[0] == '-' {
			for i := 1; i < len(arg); i++ {
				if arg[i] >= 'a' && arg[i] <= 'z' {
					options |= 1 << (arg[i] - 'a')
				} else {
					fmt.Println("Invalid Option")
					return
				}
			}
		} else {
			fmt.Println("Invalid Option")
			return
		}
	}

	fmt.Print("Options: ")

	// Iterate over each bit position from left to right (highest to lowest)
	for i := 25; i >= 0; i-- {
		bit := (options >> i) & 1
		if i == 17 || i == 9 || i == 1 { // Add space separators
			fmt.Print(" ")
		}
		if bit == 1 {
			fmt.Print(string('z' - i))
		} else {
			fmt.Print("*")
		}
	}

	fmt.Println()
}
