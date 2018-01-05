package Problem0488

var maxCount = 6

func findMinStep(board, hand string) int {
	handCount := make([]int, 26)
	for i := 0; i < len(hand); i++ {
		handCount[hand[i]-'A']++
	}

	res := helper(board+"#", handCount)
	if res == maxCount {
		return -1
	}
	return res
}

func helper(board string, handCount []int) int {
	board = removeBalls(board)
	if board == "#" {
		return 0
	}
	res, need := maxCount, 0

	for i, j := 0, 0; j < len(board); j++ {
		if board[j] == board[i] {
			continue
		}
		need = 3 - (j - i)
		if handCount[board[i]-'A'] >= need {
			handCount[board[i]-'A'] -= need
			res = min(res, need+helper(board[:i]+board[j:], handCount))
			handCount[board[i]-'A'] += need
		}
		i = j
	}
	return res
}

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
