package Problem0678

func checkValidString(s string) bool {
	min, max := 0, 0
	for i := 0; i < len(s); i++ {

		if s[i] == '(' {
			min++
			max++
		} else if s[i] == ')' {
			min--
			max--
			if max < 0 {
				return false
			}
		} else {
			min--
			max++
		}
	}

	return min <= 0 && 0 <= max
}
