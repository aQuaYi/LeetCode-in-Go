package problem1044

func longestDupSubstring(S string) string {
	res := ""

	isExist := func(L int) bool {
		seen := make(map[string]bool, len(S)-L+2)
		for i := 0; i+L <= len(S); i++ {
			sub := S[i : i+L]
			if seen[sub] {
				res = sub
				return true
			}
			seen[sub] = true
		}
		return false
	}

	lo, hi := 0, len(S)
	for lo < hi {
		mi := (lo + hi + 1) / 2
		if isExist(mi) {
			lo = mi
		} else {
			hi = mi - 1
		}
	}

	return res
}
