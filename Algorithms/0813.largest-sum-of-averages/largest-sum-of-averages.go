package problem0813

import (
	"math"
)

// dp[n][k] == 20 表示
// 长度为 n 的数组 A
// 分成 k 份的时候
// 满足题意的结果为 20
var dp [101][101]float64

func largestSumOfAverages(A []int, K int) float64 {
	n := len(A)

	dp = [101][101]float64{}

	sum := 0
	for i := range A {
		sum += A[i]
		dp[i+1][1] = float64(sum) / float64(i+1)
	}

	return search(n, K, A)
}

func search(n, k int, A []int) float64 {
	if dp[n][k] > 0 {
		return dp[n][k]
	}

	if n < k {
		return 0
	}

	sum := 0

	for i := n - 1; i > 0; i-- {
		sum += A[i]
		dp[n][k] = math.Max(dp[n][k], search(i, k-1, A)+float64(sum)/float64(n-i))
	}

	return dp[n][k]
}
