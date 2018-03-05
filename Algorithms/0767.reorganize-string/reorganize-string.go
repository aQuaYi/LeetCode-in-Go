package Problem0767

func reorganizeString(s string) string {
	bs := []byte(s)

	for i := 1; i < len(bs); i++ {
		if bs[i-1] != bs[i] {
			continue
		}

		j := i + 1
		for j < len(bs) && bs[i] == bs[j] {
			j++
		}

		if j == len(bs) {
			return ""
		}

		bs[i], bs[j] = bs[j], bs[i]
	}

	return string(bs)
}
