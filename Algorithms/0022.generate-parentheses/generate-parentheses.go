package Problem0022

func generateParenthesis(n int) []string {
	res := []string{}
	gen(n, n, "", &res)
	return res
}

func gen(left, right int, substr string, res *[]string) {
	// 成功找到一个解
	if left == 0 && right == 0 {
		*res = append(*res, substr)
		return
	}

	// 因为左括号不用担心匹配问题，只要还有左括号，就可以随便加。
	if left > 0 {
		gen(left-1, right, substr+"(", res)
	}

	// 右括号只有在左括号剩余较少的前提下，才能加
	if right > 0 && left < right {
		gen(left, right-1, substr+")", res)
	}
}
