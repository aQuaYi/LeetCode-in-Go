package problem0680

func validPalindrome(s string) bool {
	return helper([]byte(s), 0, len(s)-1, false)
}

func helper(bs []byte, lo, hi int, hasDeleted bool) bool {
	for lo < hi {
		if bs[lo] != bs[hi] {
			if hasDeleted {
				return false
			}
			return helper(bs, lo+1, hi, true) || // 删除 s[lo]
				helper(bs, lo, hi-1, true) // 删除 s[hi]
		}
		lo++
		hi--
	}

	return true
}
