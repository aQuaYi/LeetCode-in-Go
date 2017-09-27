package Problem0168

func convertToTitle(n int) string {
	if n == 0 {
		return ""
	}

	// byte('A') == 65
	str := string(byte(64 + n%26))

	return convertToTitle(n/26) + str
}
