package problem1042

import "sort"

func gardenNoAdj(N int, paths [][]int) []int {
	for _, p := range paths {
		if p[0] > p[1] {
			p[0], p[1] = p[1], p[0]
		}
	}
	sort.Slice(paths, func(i int, j int) bool {
		if paths[i][0] == paths[j][0] {
			return paths[i][1] < paths[j][1]
		}
		return paths[i][0] < paths[j][0]
	})
	res := make([]int, N)
	res[0] = 1
	for _, p := range paths {
		x, y := p[0]-1, p[1]-1
		if res[x] == 0 {
			res[x], res[y] = 1, 2
			continue
		}
		res[y] = next(res[x])
	}

	return res
}

func next(i int) int {
	if i == 4 {
		return 1
	}
	return i + 1
}
