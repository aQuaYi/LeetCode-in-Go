package problem0828

func uniqueLetterString(s string) int {
	size := len(s)
	res := 0
	for i := 0; i < size; i++ {
		for j := i + 1; j <= size; j++ {
			res += uniqueNum(s[i:j])
		}
	}
	return res
}

func uniqueNum(s string) int {
	count := map[rune]int{}
	res := 0
	for _, b := range s {
		count[b]++
		switch count[b] {
		case 1:
			res++
		case 2:
			res--
		}
	}
	return res
}
