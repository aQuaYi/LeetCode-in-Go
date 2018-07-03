package problem0848

func shiftingLetters(S string, shifts []int) string {
	size := len(S)
	a := s2is(S)
	shift := 0
	for i := size - 1; 0 <= i; i-- {
		shift += shifts[i]
		shift %= 26
		a[i] = (a[i] + shift) % 26
	}
	return is2s(a)
}

func s2is(s string) []int {
	res := make([]int, len(s))
	for i := range s {
		res[i] = int(s[i] - 'a')
	}
	return res
}

func is2s(a []int) string {
	bs := make([]byte, len(a))
	for i, n := range a {
		bs[i] = byte(n + 'a')
	}
	return string(bs)
}
