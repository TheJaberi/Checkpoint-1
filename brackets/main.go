package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	str := ""
	for _, args := range os.Args[1:] {
		for _, r := range args {
			if r == '(' || r == ')' || r == '[' || r == ']' || r == '{' || r == '}' {
				str += string(r)
			}
		}
		if IsOk(str) {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
		str = ""
	}
}

func Join(arr []string) (s string) {
	for _, a := range arr {
		s += a
	}
	return s
}

func IsOk(str string) bool {
	nstr := str
	for {
		nstr = Join(strings.Split(nstr, "{}"))
		nstr = Join(strings.Split(nstr, "()"))
		nstr = Join(strings.Split(nstr, "[]"))
		if nstr == str {
			break
		}
		str = nstr
	}
	return len(str) == 0
}
