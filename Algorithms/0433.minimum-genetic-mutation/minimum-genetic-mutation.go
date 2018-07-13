package problem0433

func minMutation(start string, end string, bank []string) int {
	isInBank := make(map[string]bool, len(bank))
	for _, g := range bank {
		isInBank[g] = true
	}

	cands := make([]string, 1, 1024)
	cands[0] = start
	res := 0

	for len(cands) > 0 {
		res++
		size := len(cands)
		for i := 0; i < size; i++ {
			cand := cands[i]
			if cand == end {
				return res
			}
			ms := mutations(cand, end)
			for _, m := range ms {
				if isInBank[m] {
					cands = append(cands, m)
					if m == end {
						return res
					}
				}
			}
		}
		cands = cands[size:]
	}
	return -1
}

func mutations(cand, end string) []string {
	cs := []byte(cand)
	res := make([]string, 0, len(end))
	for i := range end {
		if cand[i] == end[i] {
			continue
		}
		t := cs[i]
		cs[i] = end[i]
		res = append(res, string(cs))
		cs[i] = t
	}
	return res
}
