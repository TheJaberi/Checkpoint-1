package main

import (
	"os"
)

func main() {
	if len(os.Args[1:]) != 1 {
		os.Stdout.WriteString("\n")
		return
	}
	arg := os.Args[1]
	str := ""
	for i := len(arg)-1; i >= 0; i-- {
		if arg[i] != ' ' {
			break
		}
		arg = arg[:i]
	}
	for i := len(arg)-1; i >= 0; i-- {
		if arg[i] == ' ' {
			break
		}
		str = string(arg[i]) + str
	}
	os.Stdout.WriteString(str)
}
