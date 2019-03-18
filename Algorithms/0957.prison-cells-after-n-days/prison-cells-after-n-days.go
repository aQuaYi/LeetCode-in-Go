package problem0957

func prisonAfterNDays(cells []int, N int) []int {
	N = (N-1)%14 + 1 // every 14 steps is a loop

	n := cells2int(cells)

	for i := 0; i < N; i++ {
		n = (^((n << 1) ^ (n >> 1))) & 0x7e
	}

	return int2cells(n)
}

func cells2int(cells []int) int {
	res := 0
	size := len(cells)
	for i, c := range cells {
		res += c << uint(size-i-1)
	}
	return res
}

func int2cells(n int) []int {
	bits := uint(8)
	cells := make([]int, bits)
	t := bits - 1
	for i := range cells {
		cells[i] = (n >> t) & 1
		t--
	}
	return cells
}
