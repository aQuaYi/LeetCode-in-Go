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

	check := func(cap int) bool {
		if cap < maxWeight {
			return false
		}
		i, j, d := 0, 0, 0
		for ; d < D && j < n; j++ {
			if sum[j+1]-sum[i] <= cap {
				continue
			}
			i = j
			d++
		}
		return d+1 <= D && sum[n]-sum[i] <= cap
	}

	res := sort.Search(sum[n], check)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
