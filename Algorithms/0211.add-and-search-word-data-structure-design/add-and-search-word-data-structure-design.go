package problem0211

type WordDictionary struct {
	sons [26]*WordDictionary
	end  int
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	for _, b := range word {
		idx := b - 'a'
		if this.sons[idx] == nil {
			this.sons[idx] = &WordDictionary{}
		}
		this = this.sons[idx]
	}

	this.end++
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	for i, b := range word {
		if b != '.' {
			idx := b - 'a'
			if this.sons[idx] == nil {
				return false
			}

			this = this.sons[idx]
		} else {
			for _, son := range this.sons {
				if son == nil {
					continue
				}

				this = son
				if i == len(word)-1 {
					if this.end > 0 {
						return true
					}
					continue
				}

				if this.Search(word[i+1:]) {
					return true
				}
			}

			return false
		}
	}

	if this.end > 0 {
		return true
	}

	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
