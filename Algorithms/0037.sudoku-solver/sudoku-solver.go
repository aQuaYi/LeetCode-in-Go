package problem0037

func solveSudoku(board [][]byte) {
	solve(board, 0)
}

func solve(board [][]byte, k int) bool {
	if k == 81 {
		return true
	}

	r, c := k/9, k%9
	if board[r][c] != '.' {
		return solve(board, k+1)
	}

	bi, bj := r/3*3, c/3*3

	isValid := func(b byte) bool {
		for n := 0; n < 9; n++ {
			if board[r][n] == b ||
				board[n][c] == b ||
				board[bi+n/3][bj+n%3] == b {
				return false
			}
		}
		return true
	}

	for b := byte('1'); b <= '9'; b++ {
		if isValid(b) {
			board[r][c] = b
			if solve(board, k+1) {
				return true
			}
		}
	}

	board[r][c] = '.'
	return false
}
