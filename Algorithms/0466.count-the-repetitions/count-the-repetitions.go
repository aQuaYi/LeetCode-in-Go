package Problem0466

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if !isPossible(s1, s2) {
		return 0
	}

	i1, i2 := 0, 0
	len1 := len(s1)
	len2 := len(s2)
	for i1 < n1*len1 {
		if s1[i1%len1] == s2[i2%len2] {
			if i2%len2 == len2-1 && i1%len1 == len1-1 {
				return n1 * ((i2 + 1) / len2) / ((i1 + 1) / len1) / n2
			}
			i2++
		}
		i1++
	}

	return i2 / len2 / n2
}

// s1 可以重组成 s2 时，返回 true
func isPossible(s1, s2 string) bool {
	r1 := [128]bool{}
	r2 := [128]bool{}

	for i := range s1 {
		r1[s1[i]] = true
	}

	for i := range s2 {
		r2[s2[i]] = true
	}

	for i := 0; i < 128; i++ {
		if r1[i] == false && r2[i] == true {
			return false
		}
	}

	return true
}
