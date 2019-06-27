package problem1030

var dr = []int{-1, 0, 0, 1}
var dc = []int{0, -1, 1, 0}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	RC := R * C
	res := make([][]int, 1, RC)
	hasSeen := [10000]bool{}
	res[0], hasSeen[r0*100+c0] = []int{r0, c0}, true

	b, e := 0, len(res)
	for b < e {
		for i := b; i < e; i++ {
			pr, pc := res[i][0], res[i][1]
			for k := 0; k < 4; k++ {
				r, c := pr+dr[k], pc+dc[k]
				if 0 <= r && r < R &&
					0 <= c && c < C &&
					!hasSeen[r*100+c] {
					res, hasSeen[r*100+c] = append(res, []int{r, c}), true
				}
			}
		}
		b, e = e, len(res)
	}

	return res
}
