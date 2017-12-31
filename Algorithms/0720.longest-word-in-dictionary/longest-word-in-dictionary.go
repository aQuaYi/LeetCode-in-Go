package Problem0720

func longestWord(words []string) string {
	rec := make(map[string]bool, len(words))
	for _, w := range words {
		rec[w] = true
	}

	res := ""

	for _, w := range words {
		if len(res) > len(w) {
			continue
		}

		t := w[:len(w)-1]
		for rec[t] {
			t = t[:len(t)-1]
		}

		if t != "" {
			continue
		}

		if len(res) < len(w) ||
			(len(res) == len(w) && res > w) {
			res = w
		}
	}

	return res
}
