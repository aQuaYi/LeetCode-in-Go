package problem1003

func isValid(S string) bool {
	bs := []byte(S)
	stack := make([]byte, len(S))
	top := -1

	for _, b := range bs {
		top++
		stack[top] = b
		switch top {
		case 0:
			if b != 'a' {
				return false
			}
		case 1:
		default:
			if b == 'c' &&
				stack[top-1] == 'b' &&
				stack[top-2] == 'a' {
				top -= 3
			}
		}
	}

	return top == -1
}
