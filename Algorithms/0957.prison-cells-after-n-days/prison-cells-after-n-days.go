package problem0957

func prisonAfterNDays(cells []int, N int) []int {
	N = (N-1)%14 + 1 // every 14 steps is a loop

	n := cells2int(cells)

	for i := 0; i < N; i++ {
		n = (^((n >> 1) ^ (n << 1))) & 0x7e
	}

	return int2cells(n)
}

func cells2int(cells []int) int {
	res, size := 0, len(cells)
	for i := 0; i < size; i++ {
		res = res<<1 + cells[i]
	}
	return res
}

func int2cells(n int) []int {
	bits := uint(8)
	cells := make([]int, bits)
	i := 0
	for n > 0 {
		cells[i] = n & 1
		n >>= 1
		i++
	}
	for i, j := 0, int(bits-1); i < j; i, j = i+1, j-1 {
		cells[i], cells[j] = cells[j], cells[i]
	}
	return cells
}
