package problem1030

var res = [10000][]int{}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	dist := [200][][]int{}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			d := abs(r-r0) + abs(c-c0)
			dist[d] = append(dist[d], []int{r, c})
		}
	}
	begin, end := 0, 0
	for d := 0; len(dist[d]) > 0; d++ {
		begin, end = end, end+len(dist[d])
		copy(res[begin:end], dist[d])
	}
	return res[:end]
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
