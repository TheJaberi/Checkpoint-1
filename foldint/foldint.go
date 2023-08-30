package piscine

import "github.com/01-edu/z01"

func FoldInt(f func(int, int) int, a []int, n int) {
	result := n
	for _, num := range a {
		result = f(result, num)
	}
	PrintStringln(Itoa(result))
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	for n != 0 {
		s = string(n%10+'0') + s
		n /= 10
	}
	return s
}

func PrintStringln(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
