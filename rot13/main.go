package main

import (
	"os"
)

func main() {
	if len(os.Args[1:]) != 1 {
		os.Stdout.WriteString("\n")
		return
	}
	for _, r := range os.Args[1] {
		if (r > 'm' && r <= 'z') || (r > 'M' && r <= 'Z') {
			os.Stdout.WriteString(string(r-13))
		} else if (r >= 'a' && r <= 'm') || (r >= 'A' && r <= 'M') {
			os.Stdout.WriteString(string(r+13))
		} else {
			os.Stdout.WriteString(string(r))
		}
	}
}
