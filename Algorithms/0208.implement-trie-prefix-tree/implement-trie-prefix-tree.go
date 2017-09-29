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
	dict map[string]bool
}

// Constructor initialize your data structure here.
func Constructor() Trie {
	d := make(map[string]bool, 1024)
	d[""] = true
	return Trie{dict: d}
}

// Insert a word into the trie.
func (t *Trie) Insert(word string) {
	for i := 1; i <= len(word); i++ {
		t.dict[word[:i]] = true
	}
}

// Search returns true if the word is in the trie.
func (t *Trie) Search(word string) bool {
	return t.dict[word]
}

// StartsWith returns true if there is any word in the trie that starts with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	return t.dict[prefix]
}
