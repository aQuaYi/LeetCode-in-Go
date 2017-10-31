package Problem0397

func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}

	if n%2 == 0 {
		return 1 + integerReplacement(n/2)
	}

	return 1 + min(integerReplacement(n+1), integerReplacement(n-1))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
