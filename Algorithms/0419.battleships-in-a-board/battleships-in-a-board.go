package problem0419

func countBattleships(board [][]byte) int {
	if len(board) == 0 {
		return 0
	}

	m, n := len(board), len(board[0])

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' && (i == 0 || board[i-1][j] == '.') && (j == 0 || board[i][j-1] == '.') {
				// count 只有在遇到的 battleship 的左上角的 'X' 后，才会 ++
				count++
			}
		}
	}

	return count
}
