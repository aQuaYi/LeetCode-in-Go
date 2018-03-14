package problem0488

var maxCount = 6

func findMinStep(board, hand string) int {
	// ball[i] = j 表示, 手上拥有 i+'A' 颜色的球的个数
	ball := make([]int, 26)
	for i := 0; i < len(hand); i++ {
		ball[hand[i]-'A']++
	}

	res := dfs(board+"#", ball)
	if res == maxCount {
		return -1
	}
	return res
}

func dfs(s string, ball []int) int {
	s = removeBalls(s)
	if s == "#" {
		return 0
	}
	res, need := maxCount, 0

	for i, j := 0, 0; j < len(s); j++ {
		if s[j] == s[i] {
			continue
		}
		need = 3 - (j - i)
		if ball[s[i]-'A'] >= need {
			ball[s[i]-'A'] -= need
			res = min(res, need+dfs(s[:i]+s[j:], ball))
			ball[s[i]-'A'] += need
		}
		i = j
	}

	return res
}

// 删除 board 中 3 个及以上同色的球
func removeBalls(board string) string {
	for i, j := 0, 0; j < len(board); j++ {
		if board[i] == board[j] {
			continue
		}
		if i+3 <= j {
			return removeBalls(board[:i] + board[j:])
		}
		i = j
	}
	return board
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
