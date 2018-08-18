package problem0878

const mod = 1e9 + 7

func nthMagicalNumber(n, a, b int) int {
	if a > b {
		a, b = b, a
	}

	m := lcm(a, b)
	l, r := a*n/2, b*n

	for {
		med := (l + r) / 2
		count := magicalOf(med, a, b, m)
		switch {
		case count < n:
			l = med + 1
		case n < count:
			r = med - 1
		default:
			res := med - min(med%a, med%b)
			return res % mod
		}
	}
}

func magicalOf(num, a, b, m int) int {
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
