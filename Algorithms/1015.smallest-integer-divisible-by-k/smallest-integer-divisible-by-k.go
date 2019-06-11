package problem1015

func smallestRepunitDivByK(K int) int {
	if K == 1 {
		return 1
	}

	if K%2 == 0 || K%5 == 0 {
		return -1
	}

	e := 0
	m := mod(e, K)
	for m != 0 {
		e++
		m = (m + mod(e, K)) % K
	}
	return e + 1
}

// return 10^n%K
func mod(n, K int) int {
	if n == 0 {
		return 1
	}

	m := mod(n/2, K)
	m = m * m % K
	if n&1 == 1 {
		m = m * 10 % K
	}
	return m
}
