package piscine

func SwapBits(octet byte) byte {
	left := octet >> 4
	right := octet << 4
	return left | right
}
