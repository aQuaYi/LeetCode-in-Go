package problem0999

var dx = [4]int{0, 0, -1, 1}
var dy = [4]int{-1, 1, 0, 0}

func numRookCaptures(board [][]byte) int {
	res := 0
	i, j := findRook(board)
	// explore 4 direction
	for k := 0; k < 4; k++ {
		x, y := i+dx[k], j+dy[k]
		for 0 <= x && x < 8 &&
			0 <= y && y < 8 &&
			board[x][y] != 'B' {
			if board[x][y] == 'p' {
				res++
				break
			}
			x += dx[k]
			y += dy[k]
		}
	}
	return res
}

func findRook(board [][]byte) (i, j int) {
	for i = 0; i < 8; i++ {
		for j = 0; j < 8; j++ {
			if board[i][j] == 'R' {
				return
			}
		}
	}
	panic("No Rook")
}
