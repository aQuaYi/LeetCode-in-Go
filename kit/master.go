package kit

// Master 是 LeetCode 的结构体
type Master struct {
	Secret    string
	WordList  []string
	IsInWords map[string]bool
	Count     int
}

// Guess word
func (m *Master) Guess(word string) int {
	m.Count--
	if !m.IsInWords[word] {
		return -1
	}
	return matches(m.Secret, word)
}

// a,b 总是一样长的
func matches(a, b string) int {
	size := len(a)
	res := 0
	for i := 0; i < size; i++ {
		if a[i] == b[i] {
			res++
		}
	}
	return res
}

// Update 更新了 m.IsInWords
func (m *Master) Update() {
	m.IsInWords = make(map[string]bool, len(m.WordList))
	for _, w := range m.WordList {
		m.IsInWords[w] = true
	}
}
