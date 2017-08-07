package Problem0032

func longestValidParentheses(s string) int {
	max, temp := 0, 0
	left := 0

	for _, b := range s {
		switch b {
		case '(':
			left++
		default:
			if left > 0 {
				left--
				temp += 2

				if temp > max {
					max = temp
				}
			} else {
				temp = 0
				left = 0
			}
		}
	}

	return max
}
