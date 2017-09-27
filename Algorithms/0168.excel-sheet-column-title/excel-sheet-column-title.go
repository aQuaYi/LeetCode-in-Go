package Problem0168

func convertToTitle(n int) string {
	if n <= 26 {
		return string(byte(64 + n))
	}

	return convertToTitle(n/26) + convertToTitle(n%26)
}
