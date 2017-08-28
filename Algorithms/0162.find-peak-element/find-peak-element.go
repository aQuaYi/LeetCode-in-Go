package Problem0162

func findPeakElement(a []int) int {
	n := len(a)

	if n <= 1 {
		return 0
	}
	// gtl: greater than left
	gtl := make([]bool, n)
	// a[-1] == -∞
	gtl[0] = true

	// gtr: greater than right
	gtr := make([]bool, n)
	// a[l] == -∞
	gtr[n-1] = true

	for i := 1; i < n; i++ {
		gtl[i] = a[i] > a[i-1]
		gtr[n-i-1] = a[n-i-1] > a[n-i]
	}

	for i := 0; i < n; i++ {
		if gtl[i] && gtr[i] {
			return i
		}
	}

	return 0
}
