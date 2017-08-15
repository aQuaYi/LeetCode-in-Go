package Problem0051

import (
	"fmt"
)

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

	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Print(res[i][j])
			fmt.Print(" + ")
		}
		fmt.Println()
	}

	return res
}

func handlBytes(bytes [][]byte, n int) []string {
	temp := make([]string, n)
	for i := 0; i <= n; i++ {
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
	// 检查 ‘/’ 方向的对角线

	// 按行填写，检查列即可
	for i := 1; i < len(m)-1; i++ {
		if m[i][col] == 'Q' {
			return false
		}
	}

	return true
}
