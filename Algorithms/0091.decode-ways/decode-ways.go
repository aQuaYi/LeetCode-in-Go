package Problem0091

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return waysOfOne(s)
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, waysOfOne(s[0:1])

	for i := 2; i <= n; i++ {
		w1, w2 := waysOfOne(s[i-1:i]), waysOfTwo(s[i-2:i])
		switch {
		case w1 == 0 && w2 == 0:
			return 0
		case w2 == 2:
			dp[i] = dp[i-1] + dp[i-2]
		case w1 == 0: // w2 == 1
			dp[i] = dp[i-2]
		case w2 == 0: // w1 == 1
			dp[i] = dp[i-1]
		}
	}

	return dp[n]
}

func waysOfOne(s string) int {
	if s == "0" {
		return 0
	}
	return 1
}

func waysOfTwo(s string) int {
	switch {
	case '3' <= s[0] && s[0] <= '9':
		return 0
	case s[0] == '2':
		if s[1] == '7' || s[1] == '8' || s[1] == '9' {
			return 0
		}
		if s[1] == '0' {
			return 1
		}
		return 2
	case s[0] == '1':
		if s[1] == '0' {
			return 1
		}
		return 2
	default:
		// s[0] == '0'
		return 0
	}
}
