package problem0856

func scoreOfParentheses(s string) int {
	res := 0
	layers := uint(0)
	size := len(s)
	for i := 0; i < size; i++ {
		if s[i] == '(' {
			layers++
		} else {
			layers--
		}
		if s[i] == '(' && s[i+1] == ')' {
			res += 1 << (layers - 1)
		}
	}
	return res
}
