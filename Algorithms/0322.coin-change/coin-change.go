package problem0322

func coinChange(coins []int, amount int) int {
	// dp[i] 更换金额为 i 时，最小的零钱数量
	// res = dp[amount]
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, c := range coins {
			if c <= i && dp[i] > dp[i-c]+1 {
				dp[i] = dp[i-c] + 1
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

// func coinChange(coins []int, amount int) int {
// 	max := amount + 1
// 	dp := make([]int, amount+1)
// 	for i := range dp {
// 		dp[i] = max
// 	}
// 	dp[0] = 0

// 	sort.Ints(coins)

// 	for i := 1; i <= amount; i++ {
// 		for _, c := range coins {
// 			if i-c < 0 {
// 				break
// 			}
// 			dp[i] = min(dp[i], dp[i-c]+1)
// 		}
// 	}

// 	if dp[amount] > amount {
// 		return -1
// 	}

// 	return dp[amount]
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
