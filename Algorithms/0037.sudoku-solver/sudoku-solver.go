package Problem0037

func solveSudoku(board [][]byte) {
	nums := []byte("123456789")
	blocks := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	if !fillBlock(board, nums, blocks) {
		panic("此题无解")
	}
}

func fillBlock(board [][]byte, nums []byte, blocks []int) bool {
	if len(blocks) == 0 {
		return true
	}
	if len(nums) == 0 {
		return fillBlock(board, []byte("123456789"), blocks[1:])
	}

	block := blocks[0]
	n := nums[0]
	nums = nums[1:]

	// print(board, n, block)

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
		return fillBlock(board, nums, blocks)
	}
	// 检查(r,c)所在的行和列是否已经存在 n 了
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
				if fillBlock(board, nums, blocks) {
					return true
				}
				// 后面的填写不成功，所以需要还原这个格子
				board[r][c] = '.'
				// print(board, n, block)
			}
		}
	}

	return false
}
