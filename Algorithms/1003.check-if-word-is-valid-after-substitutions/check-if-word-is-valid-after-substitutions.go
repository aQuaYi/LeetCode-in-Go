package problem1003

func isValid(S string) bool {
	bs := []byte(S)
	stack := make([]byte, len(S))
	i := 0

	for _, b := range bs {
		stack[i] = b
		i++
		switch i {
		case 1, 2:
			if stack[0] != 'a' {
				return false
			}
		default:
			if stack[i-3] == 'a' &&
				stack[i-2] == 'b' &&
				stack[i-1] == 'c' {
				i -= 3
			}
		}
	}
	return i == 0
}
