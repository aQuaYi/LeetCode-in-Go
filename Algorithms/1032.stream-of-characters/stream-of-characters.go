package problem1032

// StreamChecker check letters
type StreamChecker struct {
	max    int
	trie   *trie
	stream []int8
}

// Constructor returns StreamChecker
func Constructor(words []string) StreamChecker {
	m := 0
	t := &trie{}
	for _, w := range words {
		m = max(m, len(w))
		t.insert(w)
	}
	return StreamChecker{
		max:    m,
		trie:   t,
		stream: make([]int8, 0, 1024),
	}
}

// Query returns true if letter in words
func (sc *StreamChecker) Query(letter byte) bool {
	sc.stream = append(sc.stream, int8(letter-'a'))
	n, t := len(sc.stream), sc.trie
	for i := 1; i <= sc.max && i <= n; i++ {
		index := sc.stream[n-i]
		if t.next[index] == nil {
			return false
		}
		t = t.next[index]
		if t.isWord {
			return true
		}
	}
	return false
}

type trie struct {
	next   [26]*trie
	isWord bool
}

func (t *trie) insert(word string) {
	n := len(word)
	for i := n - 1; i >= 0; i-- {
		index := int(word[i] - 'a')
		if t.next[index] == nil {
			t.next[index] = &trie{}
		}
		t = t.next[index]
		if i == 0 {
			t.isWord = true
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
