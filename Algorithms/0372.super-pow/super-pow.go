package problem0372

// (a*b) % k = ((a%k)*(b%k)) % k
func superPow(a int, b []int) int {
	base := 1337

	// return a^k mod base
	var powmod func(a, k int) int
	powmod = func(a, k int) int {
		if k == 0 {
			return 1
		}

		// k 为奇数
		if k%2 == 1 {
			return (a * powmod(a, k-1)) % base
		}

		// k 为偶数
		sub := powmod(a, k/2)
		return (sub * sub) % base
	}

	n := len(b)
	if n == 0 {
		return 1
	}

	lastB := b[n-1]
	newB := b[:n-1]

	return (powmod(superPow(a, newB), 10) * powmod(a, lastB)) % base
}
