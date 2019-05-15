package problem1009

func bitwiseComplement(N int) int {
	if N == 0 || N == 1 {
		return 1 - N
	}
	x := N - 1
	x = x | (x >> 1)
	x = x | (x >> 2)
	x = x | (x >> 4)
	x = x | (x >> 8)
	x = x | (x >> 16)
	return x ^ N
}

func bit(N int) uint {
	b := uint(0)
	for N > 0 {
		b++
		N >>= 1
	}
	return b
}
