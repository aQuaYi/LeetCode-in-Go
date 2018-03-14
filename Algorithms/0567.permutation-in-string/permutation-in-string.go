package problem0567

func checkInclusion(s1 string, s2 string) bool {
	n1 := len(s1)
	n2 := len(s2)
	if n1 > n2 {
		return false
	}

	h1 := 0
	h2 := 0
	for i := 0; i < n1; i++ {
		c1 := s1[i] - 'a'
		c2 := s2[i] - 'a'
		h1 += 1 << c1
		h2 += 1 << c2
	}

	if h1 == h2 {
		return true
	}

	for i := n1; i < n2; i++ {
		cb := s2[i-n1] - 'a'
		ce := s2[i] - 'a'
		// 利用 cb 和 ce 更新 h2
		h2 += (1 << ce) - (1 << cb)
		if h1 == h2 {
			return true
		}
	}

	return false
}
