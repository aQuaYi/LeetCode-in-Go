package problem0367

func isPerfectSquare(num int) bool {
	r := intSqrt(num)
	return r*r == num
}

func intSqrt(n int) int {
	r := n

	for r*r > n {
		r = (r + n/r) / 2
	}

	return r
}
