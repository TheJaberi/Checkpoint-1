package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args[1:]) != 1 {
		return
	}
	if os.Args[1] == "" {
		z01.PrintRune('\n')
	}
	for _, r := range os.Args[1] {
		numOfPrint := 1
		if r >= 'a' && r <= 'z' {
			numOfPrint = int(r-'a') + 1
		}
		for i := 0; i < numOfPrint; i++ {
			z01.PrintRune(r)
		}
	}
}
