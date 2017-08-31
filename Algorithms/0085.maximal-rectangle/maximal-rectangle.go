package Problem0085

func maximalRectangle(mat [][]byte) int {
	m := len(mat)
	if m == 0 {
		return 0
	}
	n := len(mat[0])
	if n == 0 {
		return 0
	}

	check := func(ax, ay, bx, by int) int {
		for i := ax; i <= bx; i++ {
			for j := ay; j <= by; j++ {
				if mat[i][j] == '0' {
					return 0
				}
			}
		}
		return (bx - ax + 1) * (by - ay + 1)
	}

	max := 0
	for ax := 0; ax < m; ax++ {
		for ay := 0; ay < n; ay++ {
			for bx := ax; bx < m; bx++ {
				for by := ay; by < n; by++ {
					tmp := check(ax, ay, bx, by)
					if max < tmp {
						max = tmp
					}
				}
			}
		}
	}

	return max
}
