package problem1178

func findNumOfValidWords(words []string, puzzles []string) []int {
	m, n := len(words), len(puzzles)

	ws := make([]word, m)
	for i, w := range words {
		ws[i] = makeWord(w)
	}

	ps := make([]word, n)
	for i, p := range puzzles {
		ps[i] = makeWord(p)
	}

	res := make([]int, n)

	for i, p := range ps {
		for _, w := range ws {
			res[i] += valid(p, w)
		}
	}

	return res
}

func valid(p, w word) int {
	if p.first&w.total != 0 &&
		p.total&w.total == w.total {
		return 1
	}
	return 0
}

type word struct {
	first, total int
}

func makeWord(w string) word {
	return word{
		first: 1 << uint(w[0]-'a'),
		total: convert(w),
	}
}

func convert(w string) int {
	res := 0
	for _, l := range w {
		res |= 1 << uint(l-'a')
	}
	return res
}
