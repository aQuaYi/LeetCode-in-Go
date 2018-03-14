package problem0368

import "sort"

func largestDivisibleSubset(a []int) []int {
	size := len(a)
	if size == 0 {
		return []int{}
	}

	sort.Ints(a)
	// 注意：下面的 a 是升序的啦

	// max 是符合题意的 largest subset 的长度
	// largest subset 中最小值是 a[idx]
	max := 1
	idx := 0

	// dp[i] 是 a[i:] 中包含 a[i] 的 largest subset 的长度
	dp := make([]int, size)
	for i := range dp {
		dp[i] = 1
	}

	// next 用来记录符合题意的 largest subset 的元素的索引值
	// largest subset 中 a[i] 的下一个元素是 a[next[i]]
	// 这样的话，由 idx 就可以得出 result 了
	next := make([]int, size)

	// 从大往小，可以检查，因为
	// 如果 a[j]%a[i]==0的话，a[next[j]]%a[i]==0 肯定成立，就无需再检查了
	for i := size - 2; 0 <= i; i-- {
		for j := size - 1; i < j; j-- {
			if a[j]%a[i] != 0 {
				continue
			}
			if dp[i] < dp[j]+1 {
				// a[i] 指向 a[j] 可以让符合题意的 subset 变大
				next[i] = j
				dp[i] = dp[j] + 1
			}
			// 出现了更大的 subset
			if max < dp[i] {
				max = dp[i]
				idx = i
			}
		}
	}

	res := make([]int, max)
	for i := range res {
		res[i] = a[idx]
		idx = next[idx]
	}

	return res
}
