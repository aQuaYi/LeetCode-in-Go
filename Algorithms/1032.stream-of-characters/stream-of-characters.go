package problem1032

// StreamChecker check letters
type StreamChecker struct {
	dic []bool
}

// Constructor returns StreamChecker
func Constructor(words []string) StreamChecker {
	dic := make([]bool, 26)
	for _, word := range words {
		for _, r := range word {
			dic[int(r-'a')] = true
		}
	}
	return StreamChecker{
		dic: dic,
	}
}

// Query returns true if letter in words
func (c *StreamChecker) Query(letter byte) bool {
	return c.dic[letter-'a']
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
