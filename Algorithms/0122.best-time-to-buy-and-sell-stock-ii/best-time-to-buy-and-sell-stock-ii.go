package problem0122

func maxProfit(prices []int) int {
	res := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}

	return res
}
