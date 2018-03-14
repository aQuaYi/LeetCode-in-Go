package problem0372

// (a*b) % k = ((a%k)*(b%k)) % k
func superPow(a int, b []int) int {
	base := 1337

	// return a^k mod base
	powmod := func(a, k int) int {
		a %= base
		res := 1

		for i := 0; i < k; i++ {
			res = (res * a) % base
		}
		return res
	}

	n := len(b)
	if n == 0 {
		return 1
	}

	lastB := b[n-1]
	newB := b[:n-1]

	return (powmod(superPow(a, newB), 10) * powmod(a, lastB)) % base
}
