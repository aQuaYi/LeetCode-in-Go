package problem1124

func longestWPI(hours []int) int {
	res := 0

	rec := make(map[int]int, len(hours)+1)
	rec[0] = -1
	rec[-1] = -1

	cur := 0
	for i, h := range hours {
		if h > 8 {
			cur++
		} else {
			cur--
		}
		res = max(res, cur)
		if j, ok := rec[cur-1]; ok {
			res = max(res, i-j)
		}
		if _, ok := rec[cur]; !ok {
			rec[cur] = i
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
