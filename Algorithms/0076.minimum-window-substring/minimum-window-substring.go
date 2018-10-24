package problem0076

func minWindow(s string, t string) string {
	size, total := len(s), len(t)

	has := [128]int{}
	need := [128]int{}
	for i := range t {
		need[t[i]]++
	}

	min := size + 1
	res := ""
	for i, j, count := 0, 0, 0; j < size; j++ {
		if has[s[j]] < need[s[j]] {
			count++
		}
		has[s[j]]++

		for need[s[i]] == 0 || has[s[i]] > need[s[i]] {
			if has[s[i]] > need[s[i]] {
				has[s[i]]--
			}
			i++
		}

		temp := j - i + 1
		if count == total && min > temp {
			min = temp
			res = s[i : j+1]
		}
	}

	return res
}
