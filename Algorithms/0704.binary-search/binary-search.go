package problem0704

func search(a []int, target int) int {
	l, r := 0, len(a)-1

	for l <= r {
		m := (l + r) / 2
		switch {
		case a[m] < target:
			l = m + 1
		case target < a[m]:
			r = m - 1
		default:
			return m
		}
	}

	return -1
}
