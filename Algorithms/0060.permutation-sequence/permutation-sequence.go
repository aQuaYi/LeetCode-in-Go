package Problem0060

func getPermutation(n int, k int) string {
	res := make([]byte, n)

	rec := make([]byte, n)
	for i := 0; i < n; i++ {
		rec[i] = byte(i) + '1'
	}

	k -= 1

	base := 1
	for i := 2; i < n; i++ {
		base *= i
	}

	for i := 0; i < n-1; i++ {
		idx := k / base
		res[i] = rec[idx]
		rec = append(rec[:idx], rec[idx+1:]...)
		k %= base
		base /= (n - i - 1)
	}
	res[n-1] = rec[0]

	return string(res)
}
