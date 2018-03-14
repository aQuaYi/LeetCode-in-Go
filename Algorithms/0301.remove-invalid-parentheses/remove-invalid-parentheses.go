package problem0301

func removeInvalidParentheses(s string) []string {
	res := []string{}
	return dfs(s, "()", 0, 0, res)
}

func dfs(s, pair string, first, last int, res []string) []string {
	var i, j, count int
	var temp string

	// s[:last] 是合法的，即对于 pair 配对来说，没有无法配对的 pair[1]
	for j = last; j < len(s); j++ {
		if s[j] == pair[0] {
			count++
		} else if s[j] == pair[1] {
			count--
		}
		if count >= 0 {
			continue
		}
		// count == -1
		// 此时 s[:j+1] 中，多了一个 pair[1]
		// 需要从 s[i:j+1]中，删除一个 pair[1]，使得删除后的 s[:j] 是合法的
		for i = first; i <= j; i++ {
			// 此时 s[:j+1] 会有两种情况
			// 情况①：
			// i   j
			// ()())()
			// 0123456
			// s[i] == pair[1] && s[i-1] != pair[1] 保证被删除的是 s[1] 或 s[3]
			// 避免了，删除 s[1] 或 s[3] 或 s[4] 后导致的重复情况
			// 情况②：
			// i     j
			// ((())))))(
			// 0123456789
			// s[i] == pair[1] && i == first 保证了在删除 s[3] 后面的递归中，
			// 被删除的是 s[4] 和 s[5]，避免了由于 s[i-1] != pair[1] 导致的遗漏
			if s[i] == pair[1] && (i == first || s[i-1] != pair[1]) {
				temp = s[:i] + s[i+1:]
				// i 是对删除元素位置的一个记录
				// 换一个 i 就是对 res 的一次划分
				res = dfs(temp, pair, i, j, res)
			}
		}

		return res
	}

	// 此时 j == len(s)，由于 s[:j] 始终都是相对于 pair 合法的
	// 因为不合法的话，就会进入 for i 循环进行删除处理。
	// 所以，此时 s 相对于 pair 就是合法的啦

	// 上面的代码是
	// 从左往右，删除了 "()" pair 中无法配对的 ")"
	// 还需要
	// 从右往左，去删除 "()" pair 中无法配对的 "("
	// 通过同时逆转 s 和 pair，可以重用前面的代码，完成上述工作
	reversed := reverse(s)

	if pair == "()" {
		return dfs(reversed, ")(", 0, 0, res)
	}

	// 此时 pair == ")("
	// 说明：
	// s 是逆转的
	// reversed 才是正常顺序的。
	// reversed 已经完成了两个方向的删除工作。就是最后的答案了。
	return append(res, reversed)
}

func reverse(s string) string {
	buf := []byte(s)
	i, j := 0, len(buf)-1
	for i < j {
		buf[i], buf[j] = buf[j], buf[i]
		i++
		j--
	}
	return string(buf)
}
