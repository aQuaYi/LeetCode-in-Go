package Problem0309

func maxProfit(prices []int) int {
	n := len(prices)
	// buy[i] 表示，按照 prices[i-1] 的价格 买 后的利润
	buy := make([]int, n+1)
	// sell[i] 表示，按照 prices[i-1] 的价格 卖 后的利润
	sell := make([]int, n+1)
	buy[1] = -1 // 无法按照 prices[0] 的价格
	for i := 2; i <= n; i++ {
		buy[i] = max(sell[i-2]-prices[i-1], buy[i-1])
		sell[i] = max(sell[i-1], buy[i-1]+prices[i-1])
	}

	return sell[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
