package Problem0127

func ladderLength(beginWord string, endWord string, words []string) int {
	// 把 words 存入字典
	// 可以利用快速地添加，删除和查找单词
	dict := make(map[string]bool, len(words))
	for i := 0; i < len(words); i++ {
		dict[words[i]] = true
	}
	// 删除 dict 中的 beginWord
	delete(dict, beginWord)

	// queue 用于存放被 trans 到的 word
	queue := make([]string, 0, len(words))

	var trans func(string) bool
	// trans 用来把 words 中 word 能够 trans 到的单词，放入 queue
	// 如果 trans 过程中，遇到了 endWord，则返回 true
	trans = func(word string) bool {
		bytes := []byte(word)
		for i := 0; i < len(bytes); i++ {
			diffLetter := bytes[i]
			// 找出仅在索引 i 上不同的单词
			for j := 0; j < 26; j++ {
				b := 'a' + byte(j)
				if b == diffLetter {
					continue
				}

				bytes[i] = b

				if dict[string(bytes)] {
					if string(bytes) == endWord {
						return true
					}
					queue = append(queue, string(bytes))
					delete(dict, string(bytes))
				}
			}
			bytes[i] = diffLetter
		}

		return false
	}

	dist := 2
	isMeetEndWord := false
	trans(beginWord)
	for len(queue) > 0 {
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			word := queue[0]
			queue = queue[1:]
			if trans(word) {
				isMeetEndWord = true
				break
			}
		}

		dist++

		if isMeetEndWord {
			return dist
		}
	}

	return 0
}
