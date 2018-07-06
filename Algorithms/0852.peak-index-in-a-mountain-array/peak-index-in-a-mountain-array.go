package problem0852

func peakIndexInMountainArray(a []int) int {
	l, r := 0, len(a)-1
	for {
		m := (l + r) / 2
		switch {
		case a[m] < a[m+1]:
			l = m // 想想看，为什么不是 l = m+1
		case a[m-1] > a[m]:
			r = m // 想想看，为什么不是 r = m-1
		default:
			return m
		}
	}
}

/**
 * 通常的二分搜索， 正确值在 [l,r] 内，
 * 当 m 不是正确值的时候，需要把 m 排除在新的 [l,r] 外
 * 所以，l = m+1 或 r = m-1，正确值依然在 [l,r] 内
 *
 * 此题中的二分搜索，正确值在 (l,r) 内
 * 当 m 不是正确值的时候，需要把 m 排除在新的 (l,r) 外
 * 所以，l = m 或 r = m，正确值依然在 (l,r) 内
 */
