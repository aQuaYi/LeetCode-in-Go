package Problem0231

func isPowerOfTwo(n int) bool {
	switch {
	case n < 1:
		return false
	case n == 1:
		return true
	case n%2 == 1:
		return false
	default:
		return isPowerOfTwo(n / 2)
	}
}
