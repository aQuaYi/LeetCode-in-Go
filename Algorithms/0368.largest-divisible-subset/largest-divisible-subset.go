package Problem0368

import "sort"

func largestDivisibleSubset(a []int) []int {
	size := len(a)
	if size == 0 {
		return []int{}
	}

	sort.Ints(a)

	max := 1
	idx := 0

	dp := make([]int, size)
	for i := range dp {
		dp[i] = 1
	}

	par := make([]int, size)
	for i := range par {
		par[i] = i
	}

	for i := size - 2; 0 <= i; i-- {
		for j := size - 1; i < j; j-- {
			if a[j]%a[i] != 0 {
				continue
			}
			if dp[i] < dp[j]+1 {
				par[i] = j
				dp[i] = dp[j] + 1
			}
			if max < dp[i] {
				max = dp[i]
				idx = i
			}
		}
	}

	res := make([]int, 0, max)
	for par[idx] != idx {
		res = append(res, a[idx])
		idx = par[idx]
	}
	res = append(res, a[idx])

	return res
}
