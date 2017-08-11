package Problem0037

func solveSudoku(board [][]byte) {
	if !fill(board, '1', 0) {
		panic("此题无解")
	}
}

func fill(board [][]byte, n byte, block int) bool {
	if block == 9 {
		// 所有的block都已经填满，成功找到了解
		return true
	}

	if n == '9'+1 {
		// 此 block 已经被填满了，去填写下一个 block
		return fill(board, '1', block+1)
	}

	// print(board, n, block)

	rowBegin := (block / 3) * 3
	colBegin := (block % 3) * 3

	for r := rowBegin; r < rowBegin+3; r++ {
		for c := colBegin; c < colBegin+3; c++ {
			if board[r][c] == n {
				// block 中已经有 n 了，无需填写
				// 去填写 n+1
				return fill(board, n+1, block)
			}
		}
	}

	// 检查 (r,c) 能否存放 n
	// 使用匿名函数，避免传递参数
	isAvaliable := func(r, c int) bool {
		// 当前位置上的字符需为 '.'
		if board[r][c] != '.' {
			return false
		}

		// (r,c) 所在的行和列不能有 n
		// 在这里就可以体现，挨个往block中填写的优势了。
		for i := 0; i < 9; i++ {
			if board[r][i] == n || board[i][c] == n {
				return false
			}
		}

		return true
	}

	for r := rowBegin; r < rowBegin+3; r++ {
		for c := colBegin; c < colBegin+3; c++ {
			if isAvaliable(r, c) {
				board[r][c] = n
				if fill(board, n+1, block) {
					return true
				}

				// 把 (r,c) 还原，把 n 移入下一个可行的位置
				board[r][c] = '.'

				// print(board, n, block)
			}
		}
	}
	// n 在此 block 中无处可放。
	// 返回 false ，让 n-1 调整位置。
	return false
}
