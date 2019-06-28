package problem1033

func numMovesStones(a, b, c int) []int {
	a, b, c = sort(a, b, c)
	minM, maxM := 2, c-a-2
	if a+1 == b {
		minM--
	}
	if b+1 == c {
		minM--
	}
	return []int{minM, maxM}
}

func sort(a, b, c int) (int, int, int) {
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	return a, b, c
}
