package Problem0676

// MagicDictionary 是一个神奇的字典
type MagicDictionary struct {
	lens    map[int]bool
	keys    map[string]bool
	changes map[string]bool
}

var replaceByte = byte('?')

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

		bs := []byte(w)
		for i := 0; i < n; i++ {
			t := bs[i]
			bs[i] = replaceByte
			md.changes[string(bs)] = true
			bs[i] = t
		}
	}
}

// Search returns if there is any word in the trie that equals to the given word after modifying exactly one character
func (md *MagicDictionary) Search(word string) bool {
	n := len(word)

	if !md.lens[n] {
		return false
	}

	bs := []byte(word)
	count := 0
	for i := 0; i < n; i++ {
		t := bs[i]
		bs[i] = replaceByte
		if md.changes[string(bs)] {
			count++
		}
		bs[i] = t
	}

	return count == 1 || count == n
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dict);
 * param_2 := obj.Search(word);
 */
