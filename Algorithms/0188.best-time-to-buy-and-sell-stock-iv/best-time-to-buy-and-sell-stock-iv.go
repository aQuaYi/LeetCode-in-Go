package Problem0188

func maxProfit(k int, prices []int) int {
	size := len(prices)
	if size <= 1 {
		return 0
	}

	profits := []int{}
	tmp := 0
	for i := 1; i < size; i++ {
		diff := prices[i] - prices[i-1]

		if tmp*diff >= 0 {
			tmp += diff
			continue
		}

		profits = append(profits, tmp)
		tmp = diff
	}
	profits = append(profits, tmp)

	res := 0

	var dfs func(int, int, int)
	dfs = func(times, idx, tmp int) {
		if times == 0 || idx >= len(profits) {
			if res < tmp {
				res = tmp
			}
			return
		}

		for i := idx; i < len(profits); i++ {
			dfs(k-1, i, tmp+max(profits[:i]))
		}
	}

	dfs(k, 0, 0)
	return res
}

func max(ps []int) int {
	max, tmp := 0, 0

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
