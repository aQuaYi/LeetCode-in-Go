package problem0518

func change(amount int, coins []int) int {
	// dp[i]表示总额为i时的方案数.
	// 转移方程: dp[i] = Σdp[i - coins[j]]; 表示 总额为i时的方案数 = 总额为i-coins[j]的方案数的加和.
	dp := make([]int, amount+1)
	// 记得初始化dp[0] = 1; 表示总额为0时方案数为1.
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ { // 从coin开始遍历，小于coin的值没有意义
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}
