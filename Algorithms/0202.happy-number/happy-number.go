package Problem0202

func isHappy(n int) bool {
	rec := make(map[int]bool, 1024)

	for !rec[n] {
		rec[n] = true
		n = trans(n)
		if n == 1 {
			return true
		}
	}

	return false
}

func trans(n int) int {
	b := 10
	res := 0
	for 0 < n {
		res += square(n % b)
		n /= b
	}
	return res
}

func square(a int) int {
	return a * a
}
