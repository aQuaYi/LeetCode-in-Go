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
	var sum1, sum2 int
	n := len(s1)
	for i := 0; i < n; i++ {
		xor ^= s1[i] ^ s2[i]
		sum1 += int(s1[i]) * int(s1[i])
		sum2 += int(s2[i]) * int(s2[i])
	}
	return xor == 0 && sum1 == sum2
}
