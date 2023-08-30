package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func primeFactors(n int) []int {
	factors := make([]int, 0)
	for i := 2; i <= n; i++ {
		for n%i == 0 && isPrime(i) {
			factors = append(factors, i)
			n /= i
		}
	}
	return factors
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num < 2 {
		return
	}

	factors := primeFactors(num)
	if len(factors) == 0 {
		return
	}

	result := ""
	for i, factor := range factors {
		if i != 0 {
			result += "*"
		}
		result += strconv.Itoa(factor)
	}

	fmt.Println(result)
}
