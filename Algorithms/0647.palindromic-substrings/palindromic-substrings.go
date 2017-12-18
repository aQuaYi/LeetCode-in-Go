package Problem0647

func countSubstrings(s string) int {
	res := len(s)
	bs := []byte(s)
	for i := 0; i < len(s); i++ {
		for j := i + 2; j <= len(s); j++ {
			if isPalindromic(bs[i:j]) {
				res++
			}
		}
	}
	return res
}

func isPalindromic(bs []byte) bool {
	i, j := 0, len(bs)-1
	for i < j {
		if bs[i] != bs[j] {
			return false
		}
		i++
		j--
	}
	return true
}
