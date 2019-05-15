package problem1009

func bitwiseComplement(N int) int {
	b := bit(N)
	return (1<<b - 1) ^ N
}

func bit(N int) uint {
	b := uint(0)
	for N > 0 {
		b++
		N >>= 1
	}
	return b
}
