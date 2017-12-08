package Problem0547

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func findCircleNum(M [][]int) int {
	res := 0
	N := len(M)

	isChecked := make([]bool, N*N)

	check := func(i, j int) {
		xs := make([]int, 1, N)
		ys := make([]int, 1, N)
		xs[0] = i
		ys[0] = j

		for len(xs) > 0 {
			i = xs[0]
			j = ys[0]
			xs = xs[1:]
			ys = ys[1:]

			isChecked[i*N+j] = true

			for k := 0; k < 4; k++ {
				x := i + dx[k]
				y := j + dy[k]

				if 0 <= x && x < N &&
					0 <= y && y <= x &&
					M[x][y] == 1 &&
					!isChecked[x*N+y] {
					xs = append(xs, x)
					ys = append(ys, y)
				}
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j <= i; j++ {
			if M[i][j] == 1 && !isChecked[i*N+j] {
				res++
				check(i, j)
			}
		}
	}

	return res
}
