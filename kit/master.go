package kit

// Master 是 LeetCode 的结构体
type Master struct {
	secret    string
	isInWords map[string]bool
	count     int
}

// Guess word
func (m *Master) Guess(word string) int {
	m.count--
	if !m.isInWords[word] {
		return -1
	}
	return matches(m.secret, word)
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

// NewMaster 创建了一个新的 *Master
func NewMaster(secret string, wordList []string, count int) *Master {
	m := &Master{
		secret:    secret,
		isInWords: make(map[string]bool, len(wordList)),
		count:     count,
	}
	for _, w := range wordList {
		m.isInWords[w] = true
	}
	return m
}
