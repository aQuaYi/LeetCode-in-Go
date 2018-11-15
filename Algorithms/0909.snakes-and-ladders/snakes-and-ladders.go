package problem0909

func snakesAndLadders(p [][]int) int {
	n := len(p)
	target := n * n

	squares := make([]int, 400)
	isChecked := [401]bool{}

	squares[0] = 1
	isChecked[1] = true

	steps := 0

	// BFS
	for len(squares) > 0 {
		steps++
		size := len(squares)
		for j := 0; j < size; j++ {
			s := squares[j]
			for i := 1; i <= 6; i++ {
				si := s + i
				x, y := location(n, si)
				if p[x][y] != -1 {
					// jump
					isChecked[si] = true
					si = p[x][y]
				}
				if si == target {
					return steps
				}
				if !isChecked[si] {
					squares = append(squares, si)
					isChecked[si] = true
				}
			}
		}
		squares = squares[size:]
	}

	return -1
}

func location(n, square int) (x, y int) {
	square--
	// 首先计算正常的位置
	x, y = square/n, square%n
	if x%2 == 1 {
		// 行号为奇数的行，需要左右翻转
		y = n - 1 - y
	}
	// 整体上下翻转
	x = n - 1 - x
	return
}
