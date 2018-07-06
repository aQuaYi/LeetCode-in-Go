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
