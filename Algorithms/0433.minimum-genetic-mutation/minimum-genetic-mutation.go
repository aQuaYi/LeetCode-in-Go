package problem0433

func minMutation(start string, end string, bank []string) int {
	if start == end {
		return 1
	}

	cands := make([]string, 1, 1024)
	cands[0] = start
	res := 0

	// 记录 bank 中的 gene 是否已经添加到 cands 中。
	// 避免重复添加
	isAdded := make([]bool, len(bank))

	for len(cands) > 0 {
		res++
		size := len(cands)
		for i := 0; i < size; i++ {
			cand := cands[i]
			for i, gene := range bank {
				if isAdded[i] || !isMutation(cand, gene) {
					continue
				}
				if gene == end {
					return res
				}
				cands = append(cands, gene)
				isAdded[i] = true
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
