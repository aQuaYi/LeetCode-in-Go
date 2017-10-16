package Problem0322

import "sort"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	sort.Ints(coins)

	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i-c < 0 {
				break
			}
			dp[i] = min(dp[i], dp[i-c]+1)
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
