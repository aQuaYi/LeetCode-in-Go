package problem1030

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	dist := [200][][]int{}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			d := abs(r-r0) + abs(c-c0)
			dist[d] = append(dist[d], []int{r, c})
		}
	}
	res := make([][]int, 0, R*C)
	for d := 0; len(dist[d]) > 0; d++ {
		res = append(res, dist[d]...)
	}
	return res
}

func abs(n int) int {
	x := n >> 63
	return (n ^ x) - x
}
