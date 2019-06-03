package problem1000

func mergeStones(stones []int, K int) int {
	size := len(stones)
	if !isPossible(size, K) {
		return -1
	}

	dp := [31][31]int{}
	for i := 1; i <= size; i++ {
		for j := i; j < i+K && j <= size; j++ {
			dp[i][j] = dp[i][j-1] + stones[j]
		}
	}

	for i := 1; i <= size; i++ {
		for j := i + K - 1; j <= size; j += K {

		}
	}

	return dp[1][size]
}

func isPossible(size, K int) bool {
	for size >= K {
		size = size/K + size%K
	}
	return size == 1
}
