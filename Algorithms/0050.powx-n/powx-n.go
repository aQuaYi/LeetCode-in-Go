package Problem0050

func myPow(x float64, n int) float64 {
	if n < 0 {
		n *= -1
		x = 1 / x
	}

	return iterPow(1, x, n)
}

func iterPow(res, x float64, n int) float64 {
	if n == 0 {
		return res
	}

	if n%2 == 1 {
		return iterPow(res*x, x, n-1)
	}

	return iterPow(res, x*x, n/2)
}
