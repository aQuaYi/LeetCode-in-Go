package problem0541

func reverseStr(s string, k int) string {
	bytes := []byte(s)

	for i := 0; i < len(s); i += 2 * k {
		j := min(i+k, len(s))
		reverse(bytes[i:j])
	}

	return string(bytes)
}

func reverse(bytes []byte) {
	i, j := 0, len(bytes)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
