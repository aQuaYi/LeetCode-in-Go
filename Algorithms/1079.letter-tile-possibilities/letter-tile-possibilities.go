package problem1079

func numTilePossibilities(tiles string) int {
	count := [26]int{}
	for _, t := range tiles {
		count[t-'A']++
	}

	var dfs func() int
	dfs = func() int {
		sum := 0
		for i := 0; i < 26; i++ {
			if count[i] == 0 {
				continue
			}
			sum++
			count[i]--
			sum += dfs()
			count[i]++
		}
		return sum
	}

	return dfs()
}
