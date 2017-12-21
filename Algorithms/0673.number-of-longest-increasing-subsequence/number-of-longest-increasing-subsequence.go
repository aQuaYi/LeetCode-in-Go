package Problem0673

func findNumberOfLIS(a []int) int {
	size := len(a)
	rec := make([]int, 0, size)

	for _, n := range a {
		isUsed := false
		for i := range rec {
			if rec[i] < n {
				rec[i] = n
				isUsed = true
			}
		}
		if !isUsed {
			rec = append(rec, n)
		}
	}

	return len(rec)
}
