package problem0676

// MagicDictionary 是一个神奇的字典
type MagicDictionary struct {
	keys []string
}

// Constructor 构建了神奇字典
func Constructor() MagicDictionary {
	return MagicDictionary{
		keys: make([]string, 0, 1024),
	}
}

// BuildDict 往字典里面添加内容
func (md *MagicDictionary) BuildDict(dict []string) {
	for _, w := range dict {
		md.keys = append(md.keys, w)
	}
}

// Search returns if there is any word in the trie that equals to the given word after modifying exactly one character
func (md *MagicDictionary) Search(word string) bool {
	n := len(word)

	for _, w := range md.keys {
		if len(w) != n || w == word {
			continue
		}
		if isChanges(w, word) {
			return true
		}
	}
	return false
}

func isChanges(w, word string) bool {
	n := len(w)
	for i := range w {
		if w[i] == word[i] {
			n--
		}
	}
	return n == 1
}
