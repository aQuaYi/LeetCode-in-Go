package Problem0211

// WordDictionary 是字典
type WordDictionary struct {
	dict []string
}

// Constructor 构建 WordDictionary
// Initialize your data structure here.
func Constructor() WordDictionary {
	return WordDictionary{}
}

// AddWord 往 WordDictionary 中添加 word
// Adds a word into the data structure.
func (d *WordDictionary) AddWord(word string) {
	d.dict = append(d.dict, word)
}

// Search 返回 true 如果 WordDictionary 中包含有 word
// Returns if the word is in the data structure.
// A word could contain the dot character '.' to represent any one letter.
func (d *WordDictionary) Search(word string) bool {
	size := len(word)
	i := 0
	for _, w := range d.dict {
		// w 和 word 的长度要一致
		if len(w) != size {
			continue
		}

		i = 0
		for ; i < size; i++ {
			if word[i] != '.' && word[i] != w[i] {
				break
			}
		}
		// 匹配完成
		if i == size {
			return true
		}
	}

	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
