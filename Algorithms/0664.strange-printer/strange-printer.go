package Problem0664

func strangePrinter(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		if i > 0 &&
			(s[i] == s[i-1] || s[i] == s[0]) {
			continue
		}
		res++
	}

	return res
}

// func strangePrinter(s string) int {
// 	n:= len(s)
// 	// dp[i][j] 表示打印 s[i:j+1] 所需的最小次数
// 	dp := make([][]int, n)
// for i:=range dp {
// 		dp[i]= make([]int, n)
// 		dp[i][i] = 1
// 	}

// 	res :=

// 	return res
// }
