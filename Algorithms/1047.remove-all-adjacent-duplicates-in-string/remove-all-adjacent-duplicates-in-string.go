package problem1047

func removeDuplicates(S string) string {
	bs := []byte(S)
	stack, top := make([]byte, len(S)), -1
	for _, b := range bs {
		if top >= 0 && stack[top] == b {
			top--
		} else {
			top++
			stack[top] = b
		}
	}
	return string(stack[:top+1])
}
