package Problem0676

// MagicDictionary 是一个神奇的字典
type MagicDictionary struct {
	lens    map[int]bool
	keys    map[string]bool
	changes map[string]bool
}

// Constructor 构建了神奇字典
func Constructor() MagicDictionary {
	return MagicDictionary{
		lens:    make(map[int]bool, 1024),
		keys:    make(map[string]bool, 1024),
		changes: make(map[string]bool, 1024),
	}
}

// BuildDict 往字典里面添加内容
func (md *MagicDictionary) BuildDict(dict []string) {
	for _, w := range dict {
		n := len(w)
		md.lens[n] = true
		md.keys[w] = true

		for i := 0; i < n; i++ {
			md.changes[w[:i]+w[i+1:]] = true
		}
	}
}

// Search returns if there is any word in the trie that equals to the given word after modifying exactly one character
func (md *MagicDictionary) Search(word string) bool {
	n := len(word)

	if !md.lens[n] {
		return false
	}

	if md.keys[word] {
		return false
	}

	for i := 0; i < n; i++ {
		if md.changes[word[:i]+word[i+1:]] {
			return true
		}
	}

	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dict);
 * param_2 := obj.Search(word);
 */
