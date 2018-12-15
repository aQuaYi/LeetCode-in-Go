package problem0020

func isValid(s string) bool {
	size := len(s)

	stack := make([]byte, size)
	top := 0

	for i := 0; i < size; i++ {
		c := s[i]
		switch c {
		case '(':
			stack[top] = c + 1 // '('+1 is ')'
			top++
		case '[', '{':
			stack[top] = c + 2
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
