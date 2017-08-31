package Problem0123

func maxProfit(prices []int) int {
	size := len(prices)
	if size == 0 {
		return 0
	}

	profile := []int{}
	tmp = prices[0]
	for i := 1; i < size; i++ {
		if prices[i-1] <= prices[i] {
			continue
		}
profile = append(profile, prices[])
		profile[i-1] = prices[i] - prices[i-1]
	}
	profile[size-1] = -1

	tmp = 0
	max1, max2 := 0, 0
	for _, p := range profile {
		if p < 0 && tmp+p <= 0 {
			switch {
			case tmp > max1:
				max2, max1 = max1, tmp
			case tmp > max2:
				max2 = tmp
			}
			tmp = 0
		} else {
			tmp += p
		}
	}

	return max1 + max2
}
