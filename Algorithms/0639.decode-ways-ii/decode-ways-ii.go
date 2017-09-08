package Problem0639

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	// dp[i] 表示 s[:i+1] 的组成方式数
	dp[0], dp[1] = 1, one(s[0:1])
	// 题目要求，返回的结果不能超过 mod
	mod := 1000000007
	dmap := getMap()

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1]*dmap[s[i-1:i]] + dp[i-2]*dmap[s[i-2:i]]
		if dp[i] == 0 {
			// 子字符串 s[:i+1] 的组成方式数为 0
			// 则，s 的组成方式数肯定也为 0
			// 这时可以提前结束
			return 0
		}
		dp[i] %= mod
	}

	return dp[n]
}

func getMap() map[string]int {
	ch := "1234567890*"
	res := make(map[string]int, len(ch)+len(ch)*len(ch))

	for i := range ch {
		s := ch[i : i+1]
		res[s] = one(s)
	}
	for i := range ch {
		for j := range ch {
			s := ch[i:i+1] + ch[j:j+1]
			res[s] = two(s)
		}
	}

	return res
}

// 检查 s 是否为合格的单字符
// len(s) == 1
// 合格，  返回 1
// 不合格， 返回 0
// 注意：
//     题目保证了 s 只含有数字和'*'
func one(s string) int {
	switch s[0] {
	case '*':
		// '*' 可以表示 1~9
		return 9
	case '0':
		return 0
	default:
		return 1
	}
}

// 检查 s 是否为合格的双字符
// len(s) == 2
// 合格，  返回 1
// 不合格， 返回 0
// 注意：
//     题目保证了 s 只含有数字和'*'
func two(s string) int {
	switch s[0] {
	case '*':
		if s[1] == '*' {
			return 15
		}
		if '0' <= s[1] && s[1] <= '6' {
			return 2
		}
		return 1
	case '1':
		if s[1] == '*' {
			return 9
		}
		return 1
	case '2':
		if s[1] == '*' {
			return 6
		}
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
