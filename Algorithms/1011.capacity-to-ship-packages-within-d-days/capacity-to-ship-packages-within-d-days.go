package problem1011

import "sort"

func shipWithinDays(weights []int, D int) int {
	n := len(weights)
	sum := make([]int, n+1)
	maxWeight := 0
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + weights[i-1]
		maxWeight = max(maxWeight, weights[i-1])
	}

	check := func(capacity int) bool {
		if capacity < maxWeight {
			return false
		}
		i, j, d := 0, 1, 0
		for ; d < D && j <= n; j++ {
			if sum[j]-sum[i] <= capacity {
				continue
			}
			i = j - 1
			d++
		}
		return d < D
	}

	return sort.Search(sum[n], check)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
