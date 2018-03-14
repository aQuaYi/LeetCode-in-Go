package problem0693

func hasAlternatingBits(n int) bool {
	std := n & 3
	if std != 1 && std != 2 {
		return false
	}

	for n > 0 {
		if n&3 != std {
			return false
		}
		n >>= 2
	}

	return true
}
