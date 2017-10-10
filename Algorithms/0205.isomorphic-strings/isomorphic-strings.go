package Problem0205

func isIsomorphic(s1, s2 string) bool {
	t := make(map[byte]byte, 256)

	for i := range s1 {

		if s1[i] == s2[i] {
			return false
		}

		b2, ok := t[s1[i]]
		if ok {
			if b2 != s2[i] {
				return false
			}
		} else {
			t[s1[i]] = s2[i]
		}
	}

	return true
}
