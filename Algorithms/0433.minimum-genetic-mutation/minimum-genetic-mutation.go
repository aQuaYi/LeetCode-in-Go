package problem0433

func minMutation(start string, end string, bank []string) int {
	if start == end {
		return 1
	}

	isInBank := make(map[string]bool, len(bank))
	for _, gene := range bank {
		isInBank[gene] = true
	}

	cands := make([]string, 1, 1024)
	cands[0] = start
	res := 0

	for len(cands) > 0 {
		res++
		size := len(cands)
		for i := 0; i < size; i++ {
			cand := cands[i]
			for gene := range isInBank {
				if isMutation(cand, gene) {
					if gene == end {
						return res
					}
					cands = append(cands, gene)
					delete(isInBank, gene)
				}
			}
		}
		cands = cands[size:]
	}
	return -1
}

func isMutation(cand, g string) bool {
	count := 0
	size := len(g)
	i := 0
	for count < 2 && i < size {
		if cand[i] != g[i] {
			count++
		}
		i++
	}
	return count == 1
}
