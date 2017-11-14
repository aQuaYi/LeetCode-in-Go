package Problem0459

func repeatedSubstringPattern(s string) bool {
	primes := getPrimes(len(s) / 2)

	for _, p := range primes {
		if len(s)%p != 0 {
			continue
		}

		Len := len(s) / p
		i := Len * 2
		for i <= len(s) {
			if s[i-Len*2:i-Len] != s[i-Len:i] {
				break
			}

			i += Len
		}

		if i == len(s)+Len {
			return true
		}
	}

	return false
}

// 返回 <=n 的所有 质数
func getPrimes(n int) []int {
	res := make([]int, 0, n)
	if n < 2 {
		return res
	}

	res = append(res, 2)
	if n == 2 {
		return res
	}

	res = append(res, 3)
	if n == 3 {
		return res
	}

	for i := 3 + 2; i <= n; i += 2 {
		if isPrime(i, res) {
			res = append(res, i)
		}
	}

	return res
}

func isPrime(n int, primes []int) bool {
	for _, p := range primes {
		if n%p == 0 {
			return false
		}
	}
	return true
}
