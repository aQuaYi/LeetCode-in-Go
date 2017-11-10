package Problem0472

func findAllConcatenatedWordsInADict(words []string) []string {
	dict := make(map[int]map[string]bool)
	for _, word := range words {
		if _, ok := dict[len(word)]; !ok {
			dict[len(word)] = make(map[string]bool)
		}
		dict[len(word)][word] = true
	}

	res := make([]string, 0, len(words))
	for _, word := range words {
		if isConcatenated(word, dict, true) {
			res = append(res, word)
		}
	}

	return res
}

func isConcatenated(word string, dict map[int]map[string]bool, isWordComplete bool) bool {
	for wLen := 1; wLen < len(word); wLen++ {
		wordMap, ok := dict[wLen]
		if ok &&
			wordMap[word[:wLen]] &&
			isConcatenated(word[wLen:], dict, false) {
			return true
		}
	}

	// 单词是完整的时候，会匹配到自己, 需要剔除这种情况
	return !isWordComplete &&
		dict[len(word)][word]
}
