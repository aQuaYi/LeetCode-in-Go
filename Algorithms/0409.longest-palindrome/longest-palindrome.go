package Problem0409

func longestPalindrome(s string) int {
	res := 0
	a := [28]int{}
	for i := range s {
		a[s[i]]++
	}

	hasOne := 0
	for i := range a {
		if a[i] == 0 {
			continue
		}
		if a[i]&1 == 0 {
			res += a[i]
		} else {
			res += a[i] - 1
			hasOne = 1
		}
	}

	return res + hasOne
}
