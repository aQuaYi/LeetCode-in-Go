package Problem0127

func ladderLength(beginWord string, endWord string, words []string) int {
	dictMap := make(map[string]byte)

	for i := 0; i < len(words); i++ {
		if _, ok := dictMap[words[i]]; !ok {
			dictMap[words[i]] = byte(1)
		}
	}

	var wq wordQueue
	dist := 2

	addNextWords(beginWord, dictMap, &wq)
	for !wq.empty() {
		wqLen := wq.size()
		for i := 0; i < wqLen; i++ {
			word := wq.popNext()
			if word == endWord {
				return dist
			}

			addNextWords(word, dictMap, &wq)
		}
		dist++
	}

	return 0
}

func addNextWords(beginWord string, dictMap map[string]byte, wq *wordQueue) {
	bytes := []byte(beginWord)
	delete(dictMap, beginWord)
	for i := 0; i < len(bytes); i++ {
		diffLetter := bytes[i]
		for j := 0; j < 26; j++ {
			b := 'a' + byte(j)
			if b == diffLetter {
				continue
			}

			bytes[i] = b
			if _, ok := dictMap[string(bytes)]; ok {
				wq.push(string(bytes))
				delete(dictMap, string(bytes))
			}
		}
		bytes[i] = diffLetter
	}
}

type wordQueue struct {
	words []string
}

func (wq *wordQueue) empty() bool {
	return len(wq.words) == 0
}

func (wq *wordQueue) popNext() string {
	if len(wq.words) == 0 {
		return ""
	}

	s := wq.words[0]
	wq.words = wq.words[1:]
	return s
}

func (wq *wordQueue) push(s string) {
	wq.words = append(wq.words, s)
}

func (wq *wordQueue) size() int {
	return len(wq.words)
}
