package Problem0204

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	primes := make([]int, 1, n/2)
	primes[0] = 2

	isPrime := func(num int) bool {
		for _, p := range primes {
			if num%p == 0 {
				return false
			}
		}
		return true
	}

	for i := 3; i < n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	return len(primes)
}
