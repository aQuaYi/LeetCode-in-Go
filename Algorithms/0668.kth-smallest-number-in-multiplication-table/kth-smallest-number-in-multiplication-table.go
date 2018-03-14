package problem0668

func findKthNumber(m int, n int, k int) int {
	// 后面的 count 函数，m 较小比较有利
	if m > n {
		m, n = n, m
	}

	lo, hi := 1, m*n
	for lo < hi {
		mid := lo + (hi-lo)/2
		c := count(mid, m, n)
		if c >= k {
			// 说明 mid 大了
			// 需要降低 hi
			hi = mid
		} else {
			// 与 hi 同理，提高 lo
			// +1 是为了避免死循环
			lo = mid + 1
		}
	}

	return hi
}

// v >= a * b , 1<=a<=m, 1<=b<=n
// count 返回 a 和 b 的所有可能组合的个数
func count(v, m, n int) int {
	c := 0
	for i := 1; i <= m; i++ {
		c += min(v/i, n)
	}
	return c
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
