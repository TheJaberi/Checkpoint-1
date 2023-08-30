package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) != 3 {
		return
	}
	num1, err1 := strconv.Atoi(os.Args[1])
	num2, err2 := strconv.Atoi(os.Args[3])
	if err1 != nil || err2 != nil {
		return
	}
	op := os.Args[2]
	if op != "+" && op != "-" && op != "%" && op != "/" && op != "*" {
		return
	}
	if op == "%" && num2 == 0 {
		fmt.Println("No modulo by 0")
		return
	} else if op == "/" && num2 == 0 {
		fmt.Println("No division by 0")
		return
	}
	res := ApplyOp(num1, num2, op)
	// to check the overflow
	if op == "+" {
		if num1 > 0 && num2 > 0 && res < 0 {
			return
		} else if num1 < 0 && num2 < 0 && res > 0 {
			return
		}
	} else if op == "-" {
		if num1 > 0 && num2 < 0 && res < 0 {
			return
		} else if num1 < 0 && num2 > 0 && res > 0 {
			return
		}
	} else if op == "*"{
		if num1 != 0 && res/num1 != num2 {
			return
		} else if num2 != 0 && res/num2 != num1 {
			return
		}
	}
	fmt.Println(res)
}

func ApplyOp(num1 int, num2 int, op string) int {
	if op == "%" {
		return num1 % num2
	} else if op == "/" {
		return num1 / num2
	} else if op == "+" {
		return num1 + num2
	} else if op == "-" {
		return num1 - num2
	} else if op == "*" {
		return num1 * num2
	}
	return -1
}
