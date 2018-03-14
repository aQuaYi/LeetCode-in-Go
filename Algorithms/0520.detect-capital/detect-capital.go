package problem0520

func detectCapitalUse(word string) bool {
	head := word[:1]
	tail := word[1:]
	if isIn(head, 'A', 'Z') {
		return isIn(tail, 'A', 'Z') || isIn(tail, 'a', 'z')
	}
	return isIn(tail, 'a', 'z')
}

func isIn(s string, first, last byte) bool {
	for i := range s {
		if !(first <= s[i] && s[i] <= last) {
			return false
		}
	}
	return true
}
