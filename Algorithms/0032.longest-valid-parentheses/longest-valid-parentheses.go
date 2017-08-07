package Problem0032

func longestValidParentheses(s string) int {
	var i, j, max, temp int
	record := make([]int, len(s))
	lens := len(s)

	for i < lens {
		for i < lens-1 && !(s[i] == '(' && s[i+1] == ')') {
			i++
		}
		j = i + 1
		for i >= 0 && s[i] == '(' && j < lens && s[j] == ')' {
			record[i], record[j] = 1, 1
			i--
			j++
		}

		i = j
	}

	for _, r := range record {
		if r == 0 {
			temp = 0
		} else {
			temp++
		}

		if temp > max {
			max = temp
		}
	}

	return max
}
