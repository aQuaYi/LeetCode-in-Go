package Problem0188

import (
	"fmt"
)

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
	rec := map[string]int{}

	max := func(begin, end int) int {
		key := fmt.Sprintf("%d-%d", begin, end)
		if t, ok := rec[key]; ok {
			return t
		}

		max, tmp := 0, 0
		for i := begin; i < end && i < len(profits); i++ {
			if tmp < 0 {
				tmp = 0
			}

			tmp += profits[i]

			if max < tmp {
				max = tmp
			}
		}

		rec[key] = max
		return max
	}

	var dfs func(int, int, int)
	dfs = func(times, idx, tmp int) {
		if times == 0 || idx >= len(profits) {
			if res < tmp {
				res = tmp
			}
			return
		}

		for i := idx + 2; i < len(profits)+2; i += 2 {

			dfs(times-1, i, tmp+max(idx, i))
		}
	}

	dfs(k, 0, 0)

	return res
}
