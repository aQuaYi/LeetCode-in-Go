package Problem0132

func minCut(s string) int {
	if isPalindrome(s) {
		return 0
	}
	
	min := len(s)
	for i := 1; i < len(s); i++ {
		if isPalindrome(s[:i]) {
			temp := minCut(s[i:])
			if min > temp {
				min = temp
			}
		}
	}
	return min + 1
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
