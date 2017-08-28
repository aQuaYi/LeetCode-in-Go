package Problem0287

func findDuplicate(a []int) int {
	l := len(a)
	rec := make([]int, l)
	for i := 0; i < l; i++ {
		rec[a[i]]++
		if rec[a[i]] == 2 {
			return a[i]
		}
	}

	return 0
}
