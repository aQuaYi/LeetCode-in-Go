package Problem0688

import (
	"math"
)

func knightProbability(n int, k int, r int, c int) float64 {
	onBoard := make(map[[2]int][][2]int, n*n)

	next := func(i, j int) [][2]int {
		pre := [2]int{i, j}
		locs := make([][2]int, 0, 8)
		var ok bool
		locs, ok = onBoard[pre]
		if ok {
			return locs
		}

		for k := 0; k < 8; k++ {
			x := dx[k] + i
			y := dy[k] + j
			if 0 <= x && x < n && 0 <= y && y < n {
				locs = append(locs, [2]int{x, y})
			}
		}

		onBoard[pre] = locs
		return locs
	}

	queue := make([][2]int, 1, 2048)
	queue[0] = [2]int{r, c}
	fm := math.Pow(8., float64(k))
	for len(queue) > 0 && k > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			locs := next(queue[i][0], queue[i][1])
			queue = append(queue, locs...)
		}
		k--
		queue = queue[size:]
	}

	return float64(len(queue)) / fm
}

var dx = []int{-2, -2, -1, 1, 2, 2, 1, -1}
var dy = []int{-1, 1, -2, -2, 1, -1, 2, 2}
