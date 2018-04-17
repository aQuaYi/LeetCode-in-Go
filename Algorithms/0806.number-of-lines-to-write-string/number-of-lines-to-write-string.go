package problem0806

func numberOfLines(widths []int, S string) []int {
	res := []int{0, 0}
	if len(S) == 0 {
		return res
	}
	res[0] = 1

	for i := 0; i < len(S); i++ {
		if res[1]+widths[S[i]-'a'] > 100 {
			res[0]++
			res[1] = widths[S[i]-'a']
		} else {
			res[1] += widths[S[i]-'a']
		}
	}

	return res
}
