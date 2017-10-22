package Problem0368

import "sort"

func largestDivisibleSubset(a []int) []int {
	size := len(a)
	if size == 0 {
		return []int{}
	}

	sort.Ints(a)
	// 注意：下面的 a 是升序的啦

	// max 是符合题意的 largest subset 的长度
	// idx 是这个 largest subset 中最小值的 index
	max := 1
	idx := 0

	// dp[i] 是 a[i:] 中包含 a[i] 的 largest subset 的长度
	dp := make([]int, size)
	for i := range dp {
		dp[i] = 1
	}

	// next[i]==j 说明
	next := make([]int, size)
	for i := range next {
		next[i] = i
	}

	for i := size - 2; 0 <= i; i-- {
		for j := size - 1; i < j; j-- {
			if a[j]%a[i] != 0 {
				continue
			}
			if dp[i] < dp[j]+1 {
				next[i] = j
				dp[i] = dp[j] + 1
			}
			if max < dp[i] {
				max = dp[i]
				idx = i
			}
		}
	}

	res := make([]int, 0, max)
	for next[idx] != idx {
		res = append(res, a[idx])
		idx = next[idx]
	}
	res = append(res, a[idx])

	return res
}
