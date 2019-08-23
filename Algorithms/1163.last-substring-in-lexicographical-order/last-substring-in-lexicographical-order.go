package problem1163

func lastSubstring(s string) string {
	last, n := 0, len(s)
	for i := 1; i < n; i++ {
		if s[i-1] == s[i] {
			// all of s[last:i+1] are the same letter.
			// 'i' has no chance to be the new 'last'
			// the reason is
			// 1. s[i+1] > s[i], 'i+1' is the new 'last'
			// 2. s[i+1] < s[i], 'last' is still 'last'
			// 3. s[i+1] = s[i], 'i+1' is the new 'i'
			//    like above, need a different letter to make decision
			continue
		}
		for l := 0; i+l < n; l++ {
			if s[last+l] < s[i+l] {
				// s[i:] is bigger than s[last:]
				// new begin is here
				last = i
				break
			}
			if s[last+l] > s[i+l] {
				// s[last:] is bigger than s[i:]
				// no need to compare continue
				break
			}
		}
	}
	return s[last:]
}
