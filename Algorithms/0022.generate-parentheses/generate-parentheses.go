package Problem0022

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	if n == 1 {
		return []string{"()"}
	}

	return gen(generateParenthesis(n - 1))
}

func gen(strs []string) []string {
	res := make([]string, 0, len(strs)*3)
	for _, s := range strs {
		res = append(res, "("+s+")")
	}
	for i, s := range strs {
		res = append(res, s+"()")
		if i != len(strs)-1 {
			res = append(res, "()"+s)
		}
	}

	return res
}
