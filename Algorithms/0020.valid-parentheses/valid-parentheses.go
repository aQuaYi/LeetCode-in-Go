package problem0020

func isValid(s string) bool {
	opposite := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := make([]rune, len(s))
	top := 0
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack[top] = opposite[c]
			top++
		case ')', ']', '}':
			if top > 0 && stack[top-1] == c {
				top--
			} else {
				return false
			}
		}
	}

	return top == 0
}
