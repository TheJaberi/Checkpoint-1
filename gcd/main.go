package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) != 2 {
		return
	}
	a, erra := strconv.Atoi(os.Args[1])
	b, errb := strconv.Atoi(os.Args[2])
	if erra != nil || errb != nil || a < 0 || b < 0 {
		return
	}
	// if a > b {
	// 	a, b = b, a
	// }
	// gcd := 1
	// for i := a; i > 0; i-- {
	// 	if a%i == 0 && b%i == 0 {
	// 		gcd = i
	// 		break
	// 	}
	// }
	// fmt.Println(gcd)
	fmt.Println(gcd(a,b))
}

// or
func gcd(a,b int) int {
	for b!= 0 {
		a,b =b,a%b
	}
	return a
}
