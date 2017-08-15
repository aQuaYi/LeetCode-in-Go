package Problem0051

func solveNQueens(n int) [][]string {
	res := [][]string{}

	m := make([][]byte, n)
	for i := 0; i < n; i++ {
		m[i] = make([]byte, n)
		for j := 0; j < len(m[i]); j++ {
			m[i][j] = '.'
		}
	}

	solve(m, 0, &res)

	return res
}

func handlBytes(bytes [][]byte, n int) []string {
	temp := make([]string, n)
	for i := 0; i < n; i++ {
		temp[i] = string(bytes[i])
	}

	return temp
}

func solve(m [][]byte, row int, res *[][]string) {

	if row == len(m) {
		*res = append(*res, handlBytes(m, len(m)))
		return
	}

	for i := 0; i < len(m); i++ {
		if isAvaliable(m, row, i) {
			m[row][i] = 'Q'
			solve(m, row+1, res)
			m[row][i] = '.'
		}
	}
}

func isAvaliable(m [][]byte, row, col int) bool {
	var i, j int
	n := len(m)

	// 检查 ‘/’ 方向的对角线
	i, j = row, col
	for i > 0 && j < n-1 {
		i--
		j++
	}
	for i < n && j >= 0 {
		if m[i][j] == 'Q' {
			return false
		}
		i++
		j--
	}

	// 检查 ‘\’ 方向的对角线
	i, j = row, col
	for i > 0 && j > 0 {
		i--
		j--
	}
	for i < n && j < n {
		if m[i][j] == 'Q' {
			return false
		}
		i++
		j++
	}

	// 按行填写，检查列即可
	i, j = 0, 0
	for i < len(m) && j < len(m) {
		if m[row][j] == 'Q' || m[i][col] == 'Q' {
			return false
		}
		i++
		j++
	}

	return true
}
