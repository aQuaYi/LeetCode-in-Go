package problem0856

func scoreOfParentheses(s string) int {
	res := 0
	factor := 1
	size := len(s)
	for i := 0; i < size; i++ {
		if s[i] == '(' {
			factor *= 2
		} else {
			factor /= 2
		}
		if s[i] == '(' && s[i+1] == ')' {
			res += factor / 2
		}
	}
	return res
}
