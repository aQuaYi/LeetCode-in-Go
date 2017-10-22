package Problem0368

import "sort"

func largestDivisibleSubset(a []int) []int {
	size := len(a)
	res := make([]int, 0, size)
	if size == 0 {
		return res
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	curMax := 1
	k := 0

	dp := make([]int, size)
	for i := range dp {
		dp[i] = 1
	}

	par := make([]int, size)
	for i := range par {
		par[i] = i
	}

	for i := 1; i < size; i++ {
		for j := 0; j < i; j++ {
			if a[j]%a[i] != 0 {
				continue
			}
			if dp[i] < dp[j]+1 {
				par[i] = j
				dp[i] = dp[j] + 1
			}
			if dp[i] > curMax {
				curMax = dp[i]
				k = i
			}
		}
	}

	for par[k] != k {
		res = append(res, a[k])
		k = par[k]
	}
	res = append(res, a[k])

	return res
}
