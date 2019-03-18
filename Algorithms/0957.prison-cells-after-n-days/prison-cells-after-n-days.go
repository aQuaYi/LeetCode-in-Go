package problem0957

func prisonAfterNDays(cells []int, N int) []int {
	n := cells2int(cells)

	for i := 0; i < N; i++ {
		n = (^((n << 1) ^ (n >> 1))) & 0x7e
	}

	return int2cells(n)
}

func cells2int(cells []int) int {
	res := 0
	t := uint(7)
	for _, c := range cells {
		res += c << t
		t--
	}
	return res
}

func int2cells(n int) []int {
	cells := make([]int, 8)
	t := uint(7)
	for i := range cells {
		cells[i] = (n >> t) & 1
		t--
	}
	return cells
}
