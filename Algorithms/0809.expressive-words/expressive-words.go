package problem0809

func expressiveWords(S string, words []string) int {
	res := 0
	short, count := parse(S)
	for i := range words {
		res += check(words[i], short, count)
	}
	return res
}

func parse(s string) (string, []int) {
	short := s[0:1]
	count := make([]int, len(s))
	idx := 0
	for i := range s {
		if short[idx] != s[i] {
			short += s[i : i+1]
			idx++
		}
		count[idx]++
	}
	return short, count[:idx+1]
}

func check(w, short string, count []int) int {
	s, c := parse(w)
	if s != short {
		return 0
	}

	for i := range count {
		if (count[i] < 3 && c[i] != count[i]) ||
			(count[i] >= 3 && c[i] > count[i]) {
			return 0
		}
	}

	return 1
}
