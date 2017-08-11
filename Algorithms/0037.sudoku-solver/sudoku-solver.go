package Problem0037

func solveSudoku(board [][]byte) {
	nums := []byte("519783426")
	for _, n := range nums {
		if !fillBlock(board, n, 0) {
			panic("此题无解")
		}
	}
}

func fillBlock(board [][]byte, n byte, block int) bool {
	print(board, n, block)

	if block == 9 {
		return true
	}

	rowZero := (block / 3) * 3
	colZero := (block % 3) * 3

	// 检查block中，是否已经存在 b 了
	had := func() bool {
		for r := rowZero; r < rowZero+3; r++ {
			for c := colZero; c < colZero+3; c++ {
				if board[r][c] == n {
					return true
				}
			}
		}
		return false
	}

	if had() {
		return fillBlock(board, n, block+1)
	}
	// 检查(r,c)所在的行和列是否已经存在 b 了
	isClear := func(r, c int) bool {
		for i := 0; i < 9; i++ {
			if board[r][i] == n || board[i][c] == n {
				return false
			}
		}
		return true
	}

	for r := rowZero; r < rowZero+3; r++ {
		for c := colZero; c < colZero+3; c++ {
			if board[r][c] == '.' && isClear(r, c) {
				board[r][c] = n
				if fillBlock(board, n, block+1) {
					return true
				}
				// 后面的填写不成功，所以需要还原这个格子
				board[r][c] = '.'
				print(board, n, block)
			}
		}
	}

	return false
}
