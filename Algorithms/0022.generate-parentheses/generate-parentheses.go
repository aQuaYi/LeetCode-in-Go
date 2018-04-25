package problem0022

func generateParenthesis(n int) []string {
	res := []string{}
	bytes := make([]byte, 0, n*2)
	gen(n, n, 0, bytes, &res)
	return res
}

func gen(left, right, idx int, bytes []byte, res *[]string) {
	// 成功找到一个解
	if left == 0 && right == 0 {
		*res = append(*res, string(bytes))
		return
	}

	// 因为左括号不用担心匹配问题，只要还有左括号，就可以随便加。
	if left > 0 {
		bytes[idx] = '('
		gen(left-1, right, idx+1, bytes, res)
	}

	// 右括号只有在左括号剩余较少的前提下，才能加
	if right > 0 && left < right {
		bytes[idx] = ')'
		gen(left, right-1, idx+1, bytes, res)
	}
}
