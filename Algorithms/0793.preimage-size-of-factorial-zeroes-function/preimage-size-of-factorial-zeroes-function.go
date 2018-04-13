package problem0793

// https://leetcode.com/problems/preimage-size-of-factorial-zeroes-function/discuss/117821/Four-binary-search-solutions-based-on-different-ideas

func preimageSizeFZF(k int) int {
	l, r := 0, 5*(k+1)
	for l <= r {
		m := l + (r-l)/2
		km := zeros(m)
		if km < k {
			l = m + 1
		} else if km > k {
			r = m - 1
		} else {
			return 5
		}
	}
	return 0
}

// 返回 x! 尾部的 0 的个数
func zeros(x int) int {
	res := 0
	for x > 0 {
		res += x / 5
		x /= 5
	}
	return res
}
