package problem0917

func reverseOnlyLetters(S string) string {
	bs := []byte(S)

	left, right := 0, len(bs)-1
	for left < right {
		for left < right && !isLetter(bs[left]) {
			left++
		}
		for left < right && !isLetter(bs[right]) {
			right--
		}
		bs[left], bs[right] = bs[right], bs[left]
		left++
		right--
	}

	return string(bs)
}

func isLetter(b byte) bool {
	return 'a' <= b && b <= 'z' ||
		'A' <= b && b <= 'Z'
}
