package Problem0567

func checkInclusion(s1, s2 string) bool {
	n := len(s1)
	b1 := []byte(s1)
	b2 := []byte(s2)
	for i := 0; i+n <= len(b2); i++ {
		if isPermutation(b1, b2[i:i+n]) {
			return true
		}
	}
	return false
}

func isPermutation(s1, s2 []byte) bool {
	var xor byte
	n := len(s1)
	for i := 0; i < n; i++ {
		xor ^= s1[i]
		xor ^= s2[i]
	}
	return xor == 0
}
