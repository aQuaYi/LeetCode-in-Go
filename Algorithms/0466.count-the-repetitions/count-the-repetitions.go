package Problem0466

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	S2 := ""
	for n2 > 0 {
		S2 += s2
		n2--
	}

	i1, I2 := 0, 0
	len1 := len(s1)
	Len2 := len(S2)
	for i1 < n1*len1 {
		if s1[i1%len1] == S2[I2%Len2] {
			if I2%Len2 == Len2-1 && i1%len1 == len1-1 {
				return n1 * ((I2 + 1) / Len2) / ((i1 + 1) / len1)
			}
			I2++
		}
		i1++
	}

	return (I2 + 1) / Len2
}

// s1 可以重组成 s2 时，返回 true
func isPossible(s1, s2 string) bool {
	rec := [128]int{}
	for i := range s1 {
		rec[s1[i]]++
	}

	for i := range s2 {
		rec[s2[i]]--
	}

	for i := range rec {
		if rec[i] < 0 {
			return false
		}
	}
	return true
}
