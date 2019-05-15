package problem1009

func bitwiseComplement(N int) int {
	x := N
	x |= 1 // when x is 0
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	// if N is 000101011101
	// then  x 000111111111
	return x ^ N
}
