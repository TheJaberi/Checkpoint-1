package piscine

func ReverseBits(oct byte) byte {
	var result byte
	for i := 0; i < 8; i++ {
		result = (result << 1) | (oct & 1)
		// fmt.Printf("%08b\n", result)
		oct >>= 1
	}
	return result
}
