package Problem0289

func gameOfLife(board [][]int) {
	m := len(board)
	if m == 0 {
		return
	}

	n := len(board[0])
	if n == 0 {
		return
	}

	old := make([][]int, m)
	for i := 0; i < m; i++ {
		old[i] = make([]int, n)
		copy(old[i], board[i])
	}

	check := func(i, j int) int {
		count := 0
		for r := i - 1; r <= i+1; r++ {
			for c := j - 1; c <= j+1; c++ {
				if 0 <= r && r < m && 0 <= c && c < n &&
					!(r == i && c == j) &&
					old[r][c] == 1 {
					count++
				}
			}
		}

		if old[i][j] == 1 {
			if count < 2 || count > 3 {
				return 0
			}
			return 1
		}

		if count == 3 {
			return 1
		}

		return 0
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = check(i, j)
		}
	}
}
