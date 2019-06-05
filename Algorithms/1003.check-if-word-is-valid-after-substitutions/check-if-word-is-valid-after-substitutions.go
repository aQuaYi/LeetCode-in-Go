package problem1003

func isValid(S string) bool {
	bs := []byte(S)
	c := 1
	for c != 0 {
		bs, c = replace(bs)
	}
	return len(bs) == 0
}

func replace(bytes []byte) ([]byte, int) {
	i, c := 0, 0
	count := 0
	for c+2 < len(bytes) {
		if bytes[c] == 'a' &&
			bytes[c+1] == 'b' &&
			bytes[c+2] == 'c' {
			c += 3
			count++
		} else {
			bytes[i] = bytes[c]
			i++
			c++
		}
	}
	for c < len(bytes) {
		bytes[i] = bytes[c]
		i++
		c++
	}
	return bytes[:i], count
}
