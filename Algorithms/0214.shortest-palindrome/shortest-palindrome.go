package Problem0214

func shortestPalindrome(s string) string {
	begin := 1
	i := len(s) - 1
	for ; i >= 0; i-- {
		if s[i] == s[0] {
			l, r := 1, i-1
			for l <= r {
				if s[l] != s[r] {
					break
				}
				l++
				r--
			}
			if l > r {
				begin = i
				break
			}
		}
	}
	res := s
	for i := begin + 1; i < len(s); i++ {
		res = s[i:i+1] + res
	}

	return res
}
