package problem0091

func numDecodings(s string) int {
	n := len(s)

	dp := make([]int, n+1)
	// dp[i] 表示 s[:i+1] 的组成方式数
	dp[0], dp[1] = 1, 1
	if s[0] == '0' {
		dp[1] = 0
	}
	two, one := 0, int(s[0]-'0')
	for i := 2; i <= n; i++ {
		last := int(s[i-1] - '0')
		two, one = one*10+last, last
		dp[i] = dp[i-2]*checkTwo(two) + dp[i-1]*checkOne(one)
	}

	return dp[n]
}

func checkTwo(code int) int {
	if 10 <= code && code <= 26 {
		return 1
	}
	return 0
}

func checkOne(code int) int {
	if 1 <= code && code <= 9 {
		return 1
	}
	return 0
}

// 检查 s 是否为合格的单字符
// len(s) == 1
// 合格，  返回 1
// 不合格， 返回 0
// 注意：
//     题目保证了 s 只含有数字
func one(s string) int {
	if s == "0" {
		return 0
	}
	return 1
}

// 检查 s 是否为合格的双字符
// len(s) == 2
// 合格，  返回 1
// 不合格， 返回 0
// 注意：
//     题目保证了 s 只含有数字
func two(s string) int {
	switch s[0] {
	case '1':
		return 1
	case '2':
		if s[1] == '7' || s[1] == '8' || s[1] == '9' {
			return 0
		}
		return 1
	default:
		// '3' <= s[0] && s[0] <= '9'，或
		// s[0] == '0'
		return 0
	}
}
