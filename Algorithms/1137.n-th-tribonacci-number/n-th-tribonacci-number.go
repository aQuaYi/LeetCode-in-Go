package problem1137

func tribonacci(n int) int {
	if n < 2 {
		return n
	}
	a, b, c := 0, 1, 1
	for n > 2 {
		a, b, c = b, c, a+b+c
		n--
	}
	return c
}
