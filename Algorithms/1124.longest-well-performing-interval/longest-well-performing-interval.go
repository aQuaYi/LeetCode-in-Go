package problem1124

func longestWPI(hours []int) int {
	res := 0

	rec := make(map[int]int, len(hours)+1)

	cur := 0
	for i, h := range hours {
		if h > 8 {
			cur++
		} else {
			cur--
		}
		if cur > 0 {
			res = i + 1
		} else {
			if j, ok := rec[cur-1]; ok {
				res = max(res, i-j)
			}
			if _, ok := rec[cur]; !ok {
				rec[cur] = i
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
