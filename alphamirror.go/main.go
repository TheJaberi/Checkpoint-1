package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args[1:]) != 1 {
		z01.PrintRune('\n')
		return
	}
	for _, r := range os.Args[1] {
		if r < 'm' && r >= 'a' {
			shift := r - 'a'
			z01.PrintRune('z'-shift)
		} else if r >= 'm' && r <= 'z' {
			shift := 'z' - r
			z01.PrintRune('a'+ shift)
		} else if r < 'M' && r >= 'A' {
			shift := r - 'A'
			z01.PrintRune('Z'-shift)
		} else if r >= 'M' && r <= 'Z' {
			shift := 'Z' - r
			z01.PrintRune('A'+ shift)
		} else {
			z01.PrintRune(r)
		}
	}
}
