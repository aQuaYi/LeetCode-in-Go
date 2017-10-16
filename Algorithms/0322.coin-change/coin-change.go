package Problem0322

import (
	"fmt"
	"sort"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = 1<<31 - 1
	}

	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	fmt.Println(coins[0])

	for _, k := range coins {
		if k > amount {
			continue
		}

		if k == amount {
			return 1
		}
		i := k
		dp[i] = 1
		for i+k <= amount {
			dp[i+k] = min(dp[i+k], dp[i]+1)
			i += k
		}
		if i == k {
			break
		}
	}

	if dp[amount] == 1<<31-1 {
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
