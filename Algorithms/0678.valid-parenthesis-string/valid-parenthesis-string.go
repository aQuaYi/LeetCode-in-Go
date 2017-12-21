package Problem0678

func checkValidString(s string) bool {
	return check(0, 0, s)
}

func check(lp, i int, s string) bool {
	for i < len(s) {
		if s[i] == '(' {
			lp++
		} else if s[i] == ')' {
			lp--
		} else {
			return check(lp+1, i+1, s) ||
				check(lp, i+1, s) ||
				check(lp-1, i+1, s)
		}

		i++

		if lp < 0 {
			return false
		}
	}

	return lp == 0
}
