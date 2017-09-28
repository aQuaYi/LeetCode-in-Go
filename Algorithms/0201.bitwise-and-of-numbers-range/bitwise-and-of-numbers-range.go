package Problem0201

func rangeBitwiseAnd(m int, n int) int {
	res := m

	for i := m + 1; i <= n; i++ {
		res &= i
	}

	return res
}
