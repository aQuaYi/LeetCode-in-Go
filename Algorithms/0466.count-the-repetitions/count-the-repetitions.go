package Problem0466

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {

	if !isPossible(s1, s2) {
		return 0
	}

	S2 := ""
	for n2 > 0 {
		S2 += s2
		n2--
	}

	i1, I2 := 0, 0
	len1 := len(s1)
	for {
		if s1[i1%len1] == S2[I2] {
			I2++
		}

		if I2 == len(S2) {
			break
		}

		i1++
	}

	C2 := i1/len1 + 1

	return n1 / C2
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
