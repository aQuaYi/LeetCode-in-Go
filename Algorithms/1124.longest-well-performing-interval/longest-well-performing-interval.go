package problem1124

func longestWPI(hours []int) int {
	res, count := 0, 0
	rec := make([]int, len(hours)+2)
	for i, h := range hours {
		if h > 8 {
			count++
		} else {
			count--
		}
		//
		if count > 0 {
			res = i + 1
		} else {
			if rec[1-count] > 0 {
				res = max(res, i-rec[1-count]+1)
			}
			if rec[-count] == 0 {
				rec[-count] = i + 1
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
