package Problem0123

func maxProfit(prices []int) int {
	size := len(prices)
	if size <= 1 {
		return 0
	}

	diff := make([]int, size-1)

	for i := 1; i < size; i++ {
		diff[i-1] = prices[i] - prices[i-1]
	}

	profits := []int{}
	temp := diff[0]
	for i := 1; i < len(diff); i++ {
		if temp*diff[i] >= 0 {
			temp += diff[i]
			continue
		}

		profits = append(profits, temp)
		temp = diff[i]
	}
	profits = append(profits, temp)

	res := 0
	for i := 0; i < len(profits); i++ {
		temp = max(profits[:i]) + max(profits[i:])
		if res < temp {
			res = temp
		}
	}

	return res
}

func max(ps []int) int {
	max := 0
	tmp := 0

	for _, p := range ps {
		if tmp < 0 {
			tmp = 0
		}

		tmp += p

		if max < tmp {
			max = tmp
		}
	}

	return max
}
