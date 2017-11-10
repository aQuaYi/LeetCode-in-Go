package Problem0472

func findAllConcatenatedWordsInADict(words []string) []string {
	hasLen := make(map[int]bool, len(words))
	hasWord := make(map[string]bool, len(words))
	for _, word := range words {
		hasLen[len(word)] = true
		hasWord[word] = true
	}

	res := make([]string, 0, len(words))
	for _, word := range words {
		if isConcatenated(word, hasLen, hasWord, true) {
			res = append(res, word)
		}
	}

	return res
}

func isConcatenated(word string, hasLen map[int]bool, hasWord map[string]bool, isWordComplete bool) bool {
	for wLen := 1; wLen < len(word); wLen++ {
		if hasLen[wLen] &&
			// word[:wLen] 非常耗时，确保 wrods 存在长度为 wLen 的单词
			// 才检查 word[:wLen] 是否存在
			hasWord[word[:wLen]] &&
			isConcatenated(word[wLen:], hasLen, hasWord, false) {
			return true
		}
	}

	// 单词是完整的时候，会匹配到自己, 需要剔除这种情况
	return !isWordComplete &&
		hasLen[len(word)] &&
		hasWord[word]
}
