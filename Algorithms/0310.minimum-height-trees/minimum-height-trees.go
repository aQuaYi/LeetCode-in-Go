package Problem0310

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	eMap := make(map[int][]int, n)
	for i := 0; i < n; i++ {
		eMap[i] = make([]int, 0, n)
	}
	degrees := make(map[int]int, n)
	for _, e := range edges {
		eMap[e[0]] = append(eMap[e[0]], e[1])
		eMap[e[1]] = append(eMap[e[1]], e[0])
		degrees[e[0]]++
		degrees[e[1]]++
	}

	var leafs = make([]int, 0, n)
	var nd, leaf int

	for i := 0; i < n; i++ {
		if degrees[i] == 1 {
			leafs = append(leafs, i)
		}
	}

	for n > 2 {
		newLeafs := make([]int, 0, len(leafs))
		for _, leaf = range leafs {
			n--
			for _, nd = range eMap[leaf] {
				degrees[nd]--
				if degrees[nd] == 1 {
					newLeafs = append(newLeafs, nd)
				}
			}
		}
		leafs = newLeafs
	}

	var res = make([]int, 0, len(leafs))
	for _, v := range leafs {
		res = append(res, v)
	}

	return res
}
