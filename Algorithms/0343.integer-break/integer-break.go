package problem0343

// https://blog.csdn.net/fuxuemingzhu/article/details/80486238
func integerBreak(n int) int {
	if n == 2 {
		return 1
	}

	if n == 3 {
		return 2
	}

	res := 1
	for n > 4 {
		res *= 3
		n -= 3
	}

	return res * n
}
