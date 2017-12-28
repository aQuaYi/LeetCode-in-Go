package Problem0730

const mod int = 1e9 + 7

func helper(dp, pre, next [][]int, left int, right int) int {
	if left > right {
		return 0
	}

	if dp[left][right] != 0 {
		return dp[left][right]
	}

	total := 0

	for c := 0; c < 4; c++ {
		i := next[left][c]
		j := pre[right][c]
		if left <= i && j <= right {
			if i < j {
				total += helper(dp, pre, next, i+1, j-1) + 2
			} else if i == j {
				total++
			}
		}
	}

	total %= mod
	dp[left][right] = total

	return total
}

func countPalindromicSubsequences(s string) int {
	n := len(s)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	before := make([][]int, n)
	after := make([][]int, n)

	indexToHead := []int{n, n, n, n}
	indexToTail := []int{-1, -1, -1, -1}

	for i := 0; i < n; i++ {
		j := n - i - 1

		before[i] = make([]int, 4)
		after[j] = make([]int, 4)

		indexToHead[s[i]-'a'] = i
		indexToTail[s[j]-'a'] = j
		for c := 0; c < 4; c++ {
			before[i][c] = indexToHead[c]
			after[j][c] = indexToTail[c]
		}
	}

	return helper(dp, before, after, 0, n-1)
}
