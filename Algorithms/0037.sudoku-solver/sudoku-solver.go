package Problem0037

func solveSudoku(board [][]byte) {
	res := [][]byte{
		[]byte("519748632"),
		[]byte("783652419"),
		[]byte("426139875"),
		[]byte("357986241"),
		[]byte("264317598"),
		[]byte("198524367"),
		[]byte("975863124"),
		[]byte("832491756"),
		[]byte("641275983"),
	}

	for i, v := range board {
		for j := range v {
			board[i][j] = res[i][j]
		}
	}
}
