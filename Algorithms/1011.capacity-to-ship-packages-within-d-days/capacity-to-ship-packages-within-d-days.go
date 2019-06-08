package problem1011

func shipWithinDays(weights []int, D int) int {
	lo, hi := 0, 0
	for _, w := range weights {
		lo = max(lo, w)
		hi += w
	}

	for lo < hi {
		mid := (lo + hi) >> 1
		days, cur := 1, 0
		for _, w := range weights {
			if cur+w > mid {
				days++
				cur = 0
			}
			cur += w
		}
		if days > D {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
