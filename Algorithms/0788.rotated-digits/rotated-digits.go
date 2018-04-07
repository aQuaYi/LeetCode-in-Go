package problem0788

func rotatedDigits(n int) int {
	count := 0
	// 暴力方法，逐个检查
	for i := 2; i <= n; i++ {
		if isValid(i) {
			count++
		}
	}
	return count
}

func isValid(n int) bool {
	var hasFoundValid bool
	for n > 0 {
		switch n % 10 {
		case 2, 5, 6, 9:
			hasFoundValid = true
		case 3, 4, 7:
			return false
		}
		n /= 10
	}
	return hasFoundValid
}
