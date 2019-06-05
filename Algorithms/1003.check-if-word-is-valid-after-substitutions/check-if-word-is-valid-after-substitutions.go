package problem1003

func isValid(S string) bool {
	bs := []byte(S)
	stack := make([]byte, 0, len(S))

	for _, b := range bs {
		stack = append(stack, b)
		depth := len(stack)
		switch depth {
		case 1, 2:
			if stack[0] != 'a' {
				return false
			}
		default:
			if stack[depth-3] == 'a' &&
				stack[depth-2] == 'b' &&
				stack[depth-1] == 'c' {
				stack = stack[:depth-3]
			}
		}
	}
	return len(stack) == 0
}
