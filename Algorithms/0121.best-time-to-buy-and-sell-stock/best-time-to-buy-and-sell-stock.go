package Problem0121

func maxProfit(prices []int) int {
	max := 0
	for i, v := range prices {
		for j := i + 1; j < len(prices); j++ {
			if max < prices[j]-v {
				max = prices[j] - v
			}
		}
	}

	return max
}
