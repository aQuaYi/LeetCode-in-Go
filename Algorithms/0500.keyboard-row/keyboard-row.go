package problem0500

import (
	"strings"
)

var qRow = map[byte]bool{
	'q': true,
	'w': true,
	'e': true,
	'r': true,
	't': true,
	'y': true,
	'u': true,
	'i': true,
	'o': true,
	'p': true,
}

var aRow = map[byte]bool{
	'a': true,
	's': true,
	'd': true,
	'f': true,
	'g': true,
	'h': true,
	'j': true,
	'k': true,
	'l': true,
}

var zRow = map[byte]bool{
	'z': true,
	'x': true,
	'c': true,
	'v': true,
	'b': true,
	'n': true,
	'm': true,
}

func isAllIn(s string, isInRow map[byte]bool) bool {
	for i := range s {
		if !isInRow[s[i]] {
			return false
		}
	}
	return true
}

func findWords(words []string) []string {
	res := make([]string, 0, len(words))

	for _, word := range words {
		w := strings.ToLower(word)
		if isAllIn(w, qRow) || isAllIn(w, aRow) || isAllIn(w, zRow) {
			res = append(res, word)
		}
	}

	return res
}
