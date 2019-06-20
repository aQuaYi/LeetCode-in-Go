package problem1032

import "strings"

// StreamChecker check letters
type StreamChecker struct {
	dic [][]string
	sb  strings.Builder
}

// Constructor returns StreamChecker
func Constructor(words []string) StreamChecker {
	dic := make([][]string, 26)
	for _, word := range words {
		index := int(word[len(word)-1] - 'a')
		dic[index] = append(dic[index], word)
	}
	return StreamChecker{
		dic: dic,
	}
}

// Query returns true if letter in words
func (sc *StreamChecker) Query(letter byte) bool {
	sc.sb.WriteByte(letter)
	index := int(letter - 'a')
	words := sc.dic[index]

	return res
}
