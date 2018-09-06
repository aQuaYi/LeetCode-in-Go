package problem0891

import (
	"sort"
)

const mod = int(1e9 + 7)

func sumSubseqWidths(a []int) int {
	sort.Ints(a)

	n := len(a)
	res := 0
	times := 1
	for i := 0; i < n; i++ {
		res += (a[i] - a[n-i-1]) * times
		res %= mod
		times = (times << 1) % mod
	}

	return res
}

/**
 * https://leetcode.com/problems/sum-of-subsequence-widths/discuss/161267/C++Java1-line-Python-Sort-and-One-Pass
 * For A[i]:
 * There are i smaller numbers,
 * so there are 2 ^ i sequences in which A[i] is maximum.
 * we should do res += A[i] * (2 ^ i)
 * There are n - i - 1 bigger numbers,
 * so there are 2 ^ (n - i - 1) sequences in which A[i] is minimum.
 * we should do res -= A[i] * 2 ^ (n - i - 1)
 */
