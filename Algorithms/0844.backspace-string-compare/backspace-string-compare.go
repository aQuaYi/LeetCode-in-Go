package problem0844

func backspaceCompare(S string, T string) bool {
	return clear(S) == clear(T)
}

func clear(s string) string {
	bs := []byte(s)
	res := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if bs[i] != '#' {
			res = append(res, bs[i])
			continue
		}

		if len(res) > 0 {
			res = res[:len(res)-1]
		}
	}

	return string(res)
}
