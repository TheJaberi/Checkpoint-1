package piscine

func FoldInt(f func(int, int) int, a []int, n int) {
	result := n
	for _, num := range a {
		result = f(result,num)
	}
	fmt.Println(result)
}
