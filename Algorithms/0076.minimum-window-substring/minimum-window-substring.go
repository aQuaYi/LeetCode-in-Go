package problem0076

func minWindow(s string, t string) string {
	size, total := len(s), len(t)

	have := [128]int{}
	need := [128]int{}
	for i := range t {
		need[t[i]]++
	}

	min := size + 1
	res := ""
	for i, j, count := 0, 0, 0; j < size; j++ {
		if have[s[j]] < need[s[j]] {
			count++
		}
		have[s[j]]++

		for i < size && have[s[i]] > need[s[i]] {
			have[s[i]]--
			i++
		}

		length := j - i + 1
		if count == total && min > length {
			min = length
			res = s[i : j+1]
		}
	}

	return res
}
