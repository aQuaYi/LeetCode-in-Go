package Problem0205

func isIsomorphic(s1, s2 string) bool {
	t := make(map[byte]byte, 256)
	r := make(map[byte]bool, 256)

	for i := range s1 {
		s2i, ok := t[s1[i]]
		if ok {
			if s2i != s2[i] {
				return false
			}
		} else {
			if r[s2[i]] {
				return false
			}
			t[s1[i]] = s2[i]
			r[s2[i]] = true
		}
	}

	return true
}
