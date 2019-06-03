package problem1000

func mergeStones(stones []int, K int) int {
	size := len(stones)
	if !isPossible(size, K) {
		return -1
	}

	dp := [31][31]int{}
	for i := 1; i <= size; i++ {
		for j := i; j < i+K && j <= size; j++ {
			dp[i][j] = dp[i][j-1] + stones[j-1]
		}
	}

	round := size/K + 1
	for ro := 2; ro <= round; ro++ {
		for l, r := 1, 1+ro*(K-1); r <= size; l, r = l+1, r+1 {
			dp[l][r] = min(dp[l][r-K+1]*2+dp[r-K+2][r], dp[l][l+K-2]+2*dp[l+K-1][r])
			for k := 1; k < K; k++ {
				dp[l][r] = min(dp[l][r], dp[l][l+k-1]+dp[l+k][r-(K-k)]*2+dp[r-(K-k)+1][r])
			}
		}
	}

	return dp[1][size]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isPossible(size, K int) bool {
	for size >= K {
		size = size/K + size%K
	}
	return size == 1
}
