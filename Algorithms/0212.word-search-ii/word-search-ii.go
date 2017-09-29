package Problem0212

import "sort"

func findWords(board [][]byte, words []string) []string {
	m := len(board)
	if m == 0 {
		return nil
	}
	n := len(board[0])
	if n == 0 {
		return nil
	}

	used := make([][]bool, m)
	indexs := make([][][2]int, 26)

	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			idx := board[i][j] - 'a'
			indexs[idx] = append(indexs[idx], [2]int{i, j})
		}
	}

	var dfs func([]byte, int, int, int) bool
	// 在 board[i][j] =?= word[k]
	dfs = func(word []byte, k, i, j int) bool {
		if k == len(word) {
			return true
		}

		if i < 0 || m <= i || j < 0 || n <= j {
			return false
		}

		res := false
		if !used[i][j] && board[i][j] == word[k] {
			used[i][j] = true
			res = dfs(word, k+1, i-1, j) ||
				dfs(word, k+1, i+1, j) ||
				dfs(word, k+1, i, j-1) ||
				dfs(word, k+1, i, j+1)
			used[i][j] = false
		}

		return res
	}
	words = unique(words)
	res := make([]string, 0, len(words))
	for _, w := range words {
		ws := []byte(w)
		idx := ws[0] - 'a'
		if len(indexs[idx]) == 0 {
			continue
		}

		for _, xy := range indexs[idx] {
			i, j := xy[0], xy[1]
			if dfs(ws, 0, i, j) {
				res = append(res, w)
				break
			}
		}
	}

	sort.Strings(res)

	return res
}

func unique(words []string) []string {
	temp := make(map[string]bool, len(words))
	for _, w := range words {
		temp[w] = true
	}
	res := make([]string, 0, len(temp))
	for w := range temp {
		res = append(res, w)
	}
	return res
}
