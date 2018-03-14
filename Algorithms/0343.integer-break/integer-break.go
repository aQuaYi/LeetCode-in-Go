package problem0343

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}

	if n == 3 {
		return 2
	}

	switch n % 3 {
	case 0:
		return pow3(n / 3)
	case 1:
		return 4 * pow3(n/3-1)
	default:
		return 2 * pow3(n/3)
	}

}

func pow3(n int) int {
	if n == 0 {
		return 1
	}

	res := pow3(n >> 1)
	if n&1 == 0 {
		return res * res
	}

	return res * res * 3
}
