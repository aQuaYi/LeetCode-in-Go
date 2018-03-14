package problem0472

func findAllConcatenatedWordsInADict(words []string) []string {
	// hasLen[3] == true 表示，words 中存在长度为 3 的单词
	hasLen := make(map[int]bool, len(words))
	// hasWord["dog"] == true 表示， words 中存在单词 "dog"
	hasWord := make(map[string]bool, len(words))
	for _, word := range words {
		hasLen[len(word)] = true
		hasWord[word] = true
	}

	res := make([]string, 0, len(words))
	for _, word := range words {
		if isConcatenated(word, hasLen, hasWord, false) {
			res = append(res, word)
		}
	}

	return res
}

func isConcatenated(word string, hasLen map[int]bool, hasWord map[string]bool, isCutted bool) bool {
	for wLen := 1; wLen < len(word); wLen++ {
		if hasLen[wLen] &&
			// word[:wLen] 非常耗时，
			// 只有当 wrods 存在长度为 wLen 的单词时
			// 才检查 word[:wLen] 是否存在
			hasWord[word[:wLen]] &&
			isConcatenated(word[wLen:], hasLen, hasWord, true) {
			return true
		}
	}

	// 单词是完整的时候，会匹配到自己, 需要剔除这种情况
	// 只能对被切除过的部分，进行全单词匹配
	return isCutted &&
		hasLen[len(word)] &&
		hasWord[word]
}
