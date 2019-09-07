package problem1178

func findNumOfValidWords(words []string, puzzles []string) []int {
	count := make(map[int]int, len(words))
	for _, w := range words {
		count[mask(w)]++
	}

	res := make([]int, len(puzzles))

	for i, p := range puzzles {
		first := 1 << uint(p[0]-'a')
		pm := mask(p)
		for wm, c := range count {
			if first&wm != 0 &&
				pm&wm == wm {
				res[i] += c
			}
		}
	}

	return res
}

func mask(w string) int {
	res := 0
	for _, l := range w {
		res |= 1 << uint(l-'a')
	}
	return res
}
