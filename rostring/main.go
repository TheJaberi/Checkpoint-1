package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	spaceBefore := true
	word := ""
	arg := os.Args[1]
	for pos, r := range os.Args[1] {
		if (r >= '0' && r <= '9') || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			spaceBefore = false
			word += string(r)
		}
		if !spaceBefore && r == ' ' {
			arg = arg[pos+1:]
			break
		}
	}
	PrintStringln(TrimSpace(arg+" "+word))
}

func TrimSpace(str string)(res string) {
	word := ""
	for _, s := range str {
		if s == ' ' && word != "" {
			if len(res) == 0 {
				res += word
			} else {
				res += " " + word
			}
			word = ""
		} else if s != ' ' {
			word += string(s)
		}
	}
	if word != "" {
		if len(res) == 0 {
			res += word
		} else {
			res += " " + word
		}
	}
	return res
}

func PrintStringln(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
