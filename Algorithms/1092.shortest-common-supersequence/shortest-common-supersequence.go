package problem1092

import "strings"

func shortestCommonSupersequence(A, B string) string {
	m, n := len(A), len(B)
	// 解题思路，
	//   先求出 A 和 B 的 LCS，
	//   然后，在 LCS 上缺少的字母
	// 利用 dp 求解 LCS Z ,
	// dp[i][j]=k 表示 A[:i] 与 B[:j] 的 LCS 的长度为 k
	// 在递归过程中，会出现三种情况:
	//   1. A[i]=B[j]， 则 dp[i][j]= dp[i-1][j-1]+1
	//   2. A[i]!=B[j] 且 Z[k]!= A[i]，则
	//   3.

	dp := [1001][1001]int{}
	b := [1001][1001]int{}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				b[i][j] = 1
			} else if dp[i-1][j] >= dp[i][j-1] {
				dp[i][j] = dp[i-1][j]
				b[i][j] = 2
			} else {
				dp[i][j] = dp[i][j-1]
				b[i][j] = 3
			}
		}
	}

	var sb strings.Builder
	var super func(int, int)
	super = func(i, j int) {
		if i == 0 {
			sb.WriteString(B[:j])
			return
		}

		if j == 0 {
			sb.WriteString(A[:i])
			return
		}

		switch b[i][j] {
		case 1:
			super(i-1, j-1)
			sb.WriteByte(A[i-1])
		case 2:
			super(i-1, j)
			sb.WriteByte(A[i-1])
		case 3:
			super(i, j-1)
			sb.WriteByte(B[j-1])
		}
	}

	super(m, n)

	return sb.String()
}
