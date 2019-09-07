package problem1178

func findNumOfValidWords(words []string, puzzles []string) []int {
	count := make(map[int]int, len(words))
	for _, w := range words {
		count[mask(w)]++
	}

	res := make([]int, len(puzzles))

	for i, p := range puzzles {
		subs := subsWithHead(p)
		for _, s := range subs {
			res[i] += count[s]
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

func subsWithHead(p string) []int {
	n := len(p)
	res := make([]int, 1, 128)
	res[0] = 1 << uint(p[0]-'a')
	for i := 1; i < n; i++ {
		x := 1 << uint(p[i]-'a')
		for _, r := range res {
			res = append(res, r|x)
		}
	}
	return res
}
