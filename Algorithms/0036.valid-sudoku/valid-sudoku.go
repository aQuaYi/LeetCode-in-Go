package Problem0036

func isValidSudoku(board [][]byte) bool {
	row, column := makeIndex()

	for i := range row {
		if !check(board, row[i], column[i]) {
			return false
		}
	}
	return true
}

func check(board [][]byte, r, c [9]int) bool {
	record := make(map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		b := board[r[i]][c[i]]
		if b == '.' {
			continue
		}
		had := record[b]
		if had {
			return false
		}
		record[b] = true
	}

	return true
}

func makeIndex() (row, column [][9]int) {
	for i := 0; i < 9; i++ {
		r := [9]int{}
		c := [9]int{}
		for j := 0; j < 9; j++ {
			r[j] = i
			c[j] = j
		}
		row = append(row, r, c)
		column = append(column, c, r)
	}

	rowBase := [][9]int{
		[9]int{0, 0, 0, 1, 1, 1, 2, 2, 2},
		[9]int{3, 3, 3, 4, 4, 4, 5, 5, 5},
		[9]int{6, 6, 6, 7, 7, 7, 8, 8, 8},
	}

	columnBase := [][9]int{
		[9]int{0, 1, 2, 0, 1, 2, 0, 1, 2},
		[9]int{3, 4, 5, 3, 4, 5, 3, 4, 5},
		[9]int{6, 7, 8, 6, 7, 8, 6, 7, 8},
	}

	for i := range rowBase {
		for j := range columnBase {
			row = append(row, rowBase[i])
			column = append(column, columnBase[j])
		}
	}

	return
}
