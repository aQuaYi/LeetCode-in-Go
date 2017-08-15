package Problem0051

func solveNQueens(n int) [][]string {
	res := [][]string{}
	col := 1
	ok := false
	for col <= n {
		ans := makeTemp(n)
		col, ok = solve(ans, 1, col)
		if !ok {
			return res
		}

		res = append(res, temp)

		col++
	}

	return res
}
func handlBytes(bytes [][]byte, n int) []string {
	temp := make([]string, n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if bytes[i][j] != 'Q' {
				bytes[i][j] = '.'
			}
		}
		temp[i-1] = string(bytes[i][1 : n+1])
	}

	return temp
}

func makeTemp(n int) [][]byte {
	m := make([][]byte, n+2)
	for i := 0; i < n+2; i++ {
		m[i] = make([]byte, n+2)
	}
	return m
}

func solve(m [][]byte, row, col int) (int, bool) {
	if row == len(m)-1 {
		return 0, true
	}
	for i := col; i < len(m)-1; i++ {
		if isAvaliable(m, row, i) {
			m[row][i] = 'Q'
			_, ok := solve(m, row+1, 1)
			if ok {
				return i, true
			}
			m[row][i] = '.'
		}
	}

	return 0, false
}

func isAvaliable(m [][]byte, row, col int) bool {

	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if m[i][j] == 'Q' {
				return false
			}
		}
	}

	for i := 1; i < len(m)-1; i++ {
		if m[row][i] == 'Q' || m[i][col] == 'Q' {
			return false
		}
	}

	return true
}
