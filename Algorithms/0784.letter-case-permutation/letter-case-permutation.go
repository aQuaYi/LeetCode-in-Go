package problem0784

func letterCasePermutation(s string) []string {
	size := len(s)
	if size == 0 {
		// 递归结束
		return []string{""}
	}

	// 提取 s 的最后一个字符，作为后缀
	lastByte := s[size-1]
	postfixs := make([]string, 1, 2)
	postfixs[0] = string(lastByte)
	// 如果最后一个字符是字母的话
	// 后缀，添加其另外一种书写形式
	if b, ok := check(lastByte); ok {
		postfixs = append(postfixs, string(b))
	}

	// 利用递归，计算出 prefixs
	prefixs := letterCasePermutation(s[:size-1])

	// res 是 prefixs 和 postfixs 的乘积
	res := make([]string, 0, len(prefixs)*len(postfixs))

	for _, pre := range prefixs {
		for _, post := range postfixs {
			res = append(res, pre+post)
		}
	}

	return res
}

// 如果 b 是字母的话，
// 返回另外一种书写形式和 true
func check(b byte) (byte, bool) {
	if 'a' <= b && b <= 'z' {
		return b + 'A' - 'a', true
	}
	if 'A' <= b && b <= 'Z' {
		return b + 'a' - 'A', true
	}
	return 0, false
}
