package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("0")
		return
	}

	input, err := strconv.Atoi(os.Args[1])
	if err != nil || input <= 0 {
		fmt.Println("0")
		return
	}

	sum := 0
	for i := 2; i <= input; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}
