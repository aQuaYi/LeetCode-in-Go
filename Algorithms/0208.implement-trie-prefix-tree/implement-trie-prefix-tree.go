package problem0208

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

// Trie 是便于 word 插入与查找的数据结构
type Trie struct {
	val  byte
	sons [26]*Trie
	end  int
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	size := len(word)
	for i := 0; i < size; i++ {
		idx := word[i] - 'a'
		if node.sons[idx] == nil {
			node.sons[idx] = &Trie{val: word[i]}
		}

		node = node.sons[idx]
	}

	node.end++
}

func (this *Trie) Search(word string) bool {
	node := this
	size := len(word)
	for i := 0; i < size; i++ {
		idx := word[i] - 'a'
		if node.sons[idx] == nil {
			return false
		}
		node = node.sons[idx]
	}

	if node.end > 0 {
		return true
	}

	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	size := len(prefix)
	for i := 0; i < size; i++ {
		idx := prefix[i] - 'a'
		if node.sons[idx] == nil {
			return false
		}
		node = node.sons[idx]
	}

	return true
}
