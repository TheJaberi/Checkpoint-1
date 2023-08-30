package piscine

func findIndex(base string, char byte) int {
	for i := 0; i < len(base); i++ {
		if base[i] == char {
			return i
		}
	}
	return -1
}

func AtoiBase(s string, base string) int {
	baseLen := len(base)
	if baseLen < 2 {
		return 0
	}

	result := 0
	for i := 0; i < len(s); i++ {
		digitIndex := findIndex(base, s[i])
		if digitIndex == -1 {
			return 0 // Invalid character in string
		}
		result = result*baseLen + digitIndex
	}

	return result
}
