package problem0076

func minWindow(s string, t string) string {
	sLen, tLen := len(s), len(t)

	need := [256]int{}
	for i := range t {
		need[t[i]]++
	}

	has := [256]int{}

	min := sLen + 1
	begin, end, winBegin, winEnd, count := 0, 0, 0, 0, 0

	for ; end < sLen; end++ {
		if need[s[end]] == 0 {
			continue
		}

		if has[s[end]] < need[s[end]] {
			count++
		}
		has[s[end]]++

		if count == tLen {
			for need[s[begin]] == 0 || has[s[begin]] > need[s[begin]] {
				if has[s[begin]] > need[s[begin]] {
					has[s[begin]]--
				}
				begin++
			}

			temp := end - begin + 1
			if min > temp {
				min = temp
				winBegin = begin
				winEnd = end
			}
		}
	}

	if count < tLen {
		return ""
	}

	return s[winBegin : winEnd+1]
}
