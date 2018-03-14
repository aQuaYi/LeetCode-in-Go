package problem0188

func maxProfit(k int, prices []int) int {
	size := len(prices)
	if size <= 1 {
		return 0
	}

	if k >= size {
		return profits(prices)
	}

	local := make([]int, k+1)
	global := make([]int, k+1)

	for i := 1; i < size; i++ {
		diff := prices[i] - prices[i-1]
		for j := k; j >= 1; j-- {
			local[j] = max(global[j-1]+max(diff, 0), local[j]+diff)
			global[j] = max(local[j], global[j])
		}
	}

	return global[k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func profits(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		tmp := prices[i] - prices[i-1]
		if tmp > 0 {
			res += tmp
		}
	}

	return res
}
