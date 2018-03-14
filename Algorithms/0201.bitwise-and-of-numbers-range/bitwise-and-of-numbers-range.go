package problem0201

func rangeBitwiseAnd(m int, n int) int {
	res := uint(0)

	for m >= 1 && n >= 1 {
		if m == n {
			return m << res
		}

		m >>= 1
		n >>= 1
		res++
	}

	return 0
}
