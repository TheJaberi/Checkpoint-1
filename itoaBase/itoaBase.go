package check

func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return ""
	}

	digits := "0123456789ABCDEF"
	result := ""

	if value < 0 {
		result += "-"
		value = -value
	} else if value == 0 {
		return "0"
	}

	for value > 0 {
		digit := value % base
		result = string(digits[digit]) + result
		value /= base
	}

	return result
}
