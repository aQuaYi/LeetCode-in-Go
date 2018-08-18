package problem0878

const mod = 1e9 + 7

func nthMagicalNumber(n, a, b int) int {
	if a == b {
		return resultOfAEqualB(n, a)
	} else if a > b {
		a, b = b, a
	}

	mtp := lcm(a, b)
	l, r := a*n/2, b*n
	var med int
	for l < r {
		med = (l + r) / 2
		count := countMagicalUnderNum(med, a, b, mtp)
		if count < n {
			l = med + 1
		} else if n < count {
			r = med - 1
		} else {
			break
		}
	}

	res := med - min(med%a, med%b)

	return res % mod
}

func resultOfAEqualB(N, A int) int {
	return N * A % mod
}

func countMagicalUnderNum(num, a, b, m int) int {
	return num/a + num/b - num/m
}

// 最小公倍数
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// 最大公约数
func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
