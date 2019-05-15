package problem1009

func bitwiseComplement(N int) int {
	x := N | 1
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	return x ^ N
}
