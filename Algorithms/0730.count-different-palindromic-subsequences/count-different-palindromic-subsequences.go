package problem0730

const mod int = 1e9 + 7

// 题目中限定了，s 中只存在 a,b,c,d 这 4 个字母，在程序中，使用 0～3 替代他们
func countPalindromicSubsequences(s string) int {
	n := len(s)

	// dp[i][j] 表示 countPalindromicSubsquences(s[i:j+1])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	// before[i][0] == k 表示
	//   如果 0<=k<=i 的话， s[:i+1]中 s[k]=='a'，且 s[k+1:i+1] 中没有 'a'
	//   否则的话， s[:i+1] 中并没有 'a'
	before := make([][]int, n)
	// after[j][0] == k 表示
	//   如果 j<=k<n 的话， s[j:]中 s[k]=='a'，且 s[j:k] 中没有 'a'
	//   否则的话， s[j:] 中并没有 'a'
	after := make([][]int, n)

	// 提前设置好值，是为了 for 循环过程中，还没有出现的字母准备的
	indexOfBefore := []int{n, n, n, n}
	indexOfAfter := []int{-1, -1, -1, -1}

	for i := 0; i < n; i++ {
		j := n - i - 1

		before[i] = make([]int, 4)
		after[j] = make([]int, 4)

		indexOfBefore[s[i]-'a'] = i
		indexOfAfter[s[j]-'a'] = j
		for c := 0; c < 4; c++ {
			before[i][c] = indexOfBefore[c]
			after[j][c] = indexOfAfter[c]
		}

	}

	return helper(dp, before, after, 0, n-1)
}

func helper(dp, before, after [][]int, left int, right int) int {
	if left > right {
		return 0
	}

	// 已经计算过了，直接返回答案
	if dp[left][right] != 0 {
		return dp[left][right]
	}

	total := 0

	for c := 0; c < 4; c++ {
		i := after[left][c]
		j := before[right][c]
		if left <= i && j <= right {
			if i < j {
				// 此时
				// s[left:right+1] 中至少有 2 个 c
				// 那么可以组成的回文格式为
				//    c
				//    cc
				//    c*c，其中 * 代表 s[i+1:j] 中的不同回文数量
				// 所以，此时以 c 为两端的回文数量，等于
				//    s[i+1:j] 中的回文数量 + 2
				// 即
				//    helper(dp, before, after, i+1, j-1) + 2

				total += helper(dp, before, after, i+1, j-1) + 2
			} else if i == j {
				// i==j 说明此时只有一个单字符的回文 c
				total++
			}
		}
	}

	total %= mod
	dp[left][right] = total

	return total
}
