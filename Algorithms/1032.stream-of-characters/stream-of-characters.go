package problem1032

// StreamChecker check letters
type StreamChecker struct {
	dic    [][]string
	querys []byte
}

// Constructor returns StreamChecker
func Constructor(words []string) StreamChecker {
	dic := make([][]string, 26)
	maxLen := 0
	for _, word := range words {
		Len := len(word)
		maxLen = max(maxLen, Len)
		index := int(word[Len-1] - 'a')
		dic[index] = append(dic[index], word)
	}
	return StreamChecker{
		dic:    dic,
		querys: make([]byte, 0, 40001),
	}
}

// Query returns true if letter in words
func (sc *StreamChecker) Query(letter byte) bool {
	sc.querys = append(sc.querys, letter)
	index := int(letter - 'a')
	words := sc.dic[index]
	for _, w := range words {
		lq, lw := len(sc.querys), len(w)
		if lq-lw >= 0 && w == string(sc.querys[lq-lw:]) {
			return true
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
