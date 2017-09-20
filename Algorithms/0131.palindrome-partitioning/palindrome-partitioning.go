package Problem0131

func partition(s string) [][]string {
	res := [][]string{}
	if len(s) <= 1 {
		res = append(res, []string{s})
		return res
	}

	for i := 1; i < len(s); i++ {
		if isPalindrome(s[:i]) {
			other := partition(s[i:])
			for _, o := range other {
				temp := make([]string, len(o)+1)
				temp[0] = s[:i]
				copy(temp[1:], o)
				res = append(res, temp)
			}
		}
	}

	if isPalindrome(s) {
		res = append(res, []string{s})
	}

	return res
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
