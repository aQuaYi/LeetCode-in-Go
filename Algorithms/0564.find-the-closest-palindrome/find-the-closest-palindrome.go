package Problem0564

func nearestPalindromic(n string) string {
	if !isPalindrome(n) {
		return palindromefy(n)
	}

	return n
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func palindromefy(n string) string {
	bytes := []byte(n)
	i, j := 0, len(bytes)-1
	for i < j {
		if bytes[i] != bytes[j] {
			bytes[j] = bytes[i]
		}
		i++
		j--
	}
	return string(bytes)
}
