package problem0794

func validTicTacToe(board []string) bool {
	xs, os := count(board)

	if isWin('X', board) {
		return xs == os+1 && !isWin('O', board)
	}

	if isWin('O', board) {
		return xs == os && !isWin('X', board)
	}

	return xs == os+1 || xs == os
}

func count(board []string) (xs, os int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			switch board[i][j] {
			case 'X':
				xs++
			case 'O':
				os++
			}
		}
	}
	return
}

func isWin(c byte, b []string) bool {
	for i := 0; i < 3; i++ {
		if (b[i][0] == c && b[i][1] == c && b[i][2] == c) ||
			(b[0][i] == c && b[1][i] == c && b[2][i] == c) {
			return true
		}
	}

	return b[0][0] == c && b[1][1] == c && b[2][2] == c ||
		b[0][2] == c && b[1][1] == c && b[2][0] == c
}
