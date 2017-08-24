package Problem0079

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}

	n := len(board[0])
	if n == 0 {
		return false
	}

	if len(word) == 0 {
		return false
	}

	rec := make([][]bool, m)
	for i := range rec {
		rec[i] = make([]bool, n)
	}

	delta := [][2]int{
		[2]int{0, -1},
		[2]int{0, 1},
		[2]int{-1, 0},
		[2]int{1, 0},
	}

	var dfs func(int, int, string) bool
	dfs = func(r, c int, word string) bool {
		if len(word) == 0 {
			return true
		}

		if r < 0 || m <= r || c < 0 || n <= c {
			return false
		}

		if board[r][c] != word[0] || rec[r][c] == true {
			return false
		}

		rec[r][c] = true

		for _, d := range delta {
			if dfs(r+d[0], c+d[1], word[1:]) {
				return true
			}
		}

		rec[r][c] = false

		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, word) {
				return true
			}
		}
	}

	return false
}
