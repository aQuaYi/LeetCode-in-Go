package problem0714

func maxProfit(prices []int, fee int) int {
	empty := 0
	hold := -1 << 63

	for _, p := range prices {
		temp := empty
		empty = max(empty, hold+p)
		hold = max(hold, temp-p-fee)
	}

	return empty
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
