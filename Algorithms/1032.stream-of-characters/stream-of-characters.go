package problem1032

// StreamChecker check letters
type StreamChecker struct {
	tree   *trie
	stream []int8
	max    int
}

// Constructor returns StreamChecker
func Constructor(words []string) StreamChecker {
	maxLen := 0
	tree := &trie{}
	for _, word := range words {
		maxLen = max(maxLen, len(word))
		tree.insert(word)
	}
	return StreamChecker{
		tree:   tree,
		stream: make([]int8, 0, 40001),
		max:    maxLen,
	}
}

// Query returns true if letter in words
func (sc *StreamChecker) Query(letter byte) bool {
	sc.stream = append(sc.stream, int8(letter-'a'))
	n := len(sc.stream)
	cur := sc.tree
	for i := 1; i <= sc.max && i <= n; i++ {
		index := sc.stream[n-i]
		if cur.next[index] == nil {
			return false
		}
		cur = cur.next[index]
		if cur.isWord {
			return true
		}
	}
	return false
}

// TODO:   var isWord = &trie{}

type trie struct {
	next   [26]*trie
	isWord bool
}

func (t *trie) insert(word string) {
	n := len(word)
	cur := t
	for i := n - 1; i >= 0; i-- {
		index := int(word[i] - 'a')
		if cur.next[index] == nil {
			cur.next[index] = &trie{}
		}
		cur = cur.next[index]
		if i == 0 {
			cur.isWord = true
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
