package problem1030

import "sort"

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	d := func(point []int) int {
		return dist(point[0], point[1], r0, c0)
	}
	RC := R * C
	res := make([][]int, 0, RC)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			res = append(res, []int{i, j})
		}
	}

	sort.SliceStable(res, func(i int, j int) bool {
		return d(res[i]) < d(res[j])
	})

	return res
}

func dist(r1, c1, r2, c2 int) int {
	return abs(r1-r2) + abs(c1-c2)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
