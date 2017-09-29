package Problem0208

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

// Trie 是便于 word 插入与查找的数据结构
type Trie struct {
	words []string
}

// Constructor initialize your data structure here.
func Constructor() Trie {
	return Trie{}
}

// Insert a word into the trie.
func (t *Trie) Insert(word string) {
	t.words = append(t.words, word)
}

// Search returns true if the word is in the trie.
func (t *Trie) Search(word string) bool {
	for _, w := range t.words {
		if w == word {
			return true
		}
	}
	return false
}

// StartsWith returns true if there is any word in the trie that starts with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	end := len(prefix)
	for _, w := range t.words {
		if len(w) >= end && w[:end] == prefix {
			return true
		}
	}
	return false
}
