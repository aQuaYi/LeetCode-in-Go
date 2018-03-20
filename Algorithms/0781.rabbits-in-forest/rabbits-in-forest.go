package problem0781

func numRabbits(answers []int) int {
	isSaw := make(map[int]bool, len(answers))
	res := 0
	for _, v := range answers {
		if isSaw[v] {
			continue
		}
		res += v
		isSaw[v] = true
	}

	res += len(isSaw)

	return res
}
