package problem1033

var zerozero = []int{0, 0}

func numMovesStones(a, b, c int) []int {
	a, b, c = sort(a, b, c)
	if c-a == 2 {
		return zerozero
	}
	minM, maxM := 2, c-a-2
	if min(b-a, c-b) <= 2 {
		minM = 1
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
