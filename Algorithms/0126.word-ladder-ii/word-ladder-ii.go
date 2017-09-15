package Problem0126

func findLadders(beginWord string, endWord string, words []string) [][]string {
	res := [][]string{}

	isUsed := map[string]bool{}

	transMap := map[string][]string{}
	for i := 0; i < len(words); i++ {
		if words[i] == beginWord {
			// Note that beginWord is not a transformed word.
			continue
		}

		if isTransable(words[i], beginWord) {
			transMap[beginWord] = append(transMap[beginWord], words[i])
		}

		for j := i + 1; j < len(words); j++ {
			if words[j] == beginWord {
				// Note that beginWord is not a transformed word.
				continue
			}

			a, b := words[i], words[j]
			if isTransable(a, b) {
				transMap[a] = append(transMap[a], b)
				transMap[b] = append(transMap[b], a)
			}
		}
	}

	minLen := 1<<63 - 1

	var dfs func(string, []string, int)
	dfs = func(word string, path []string, idx int) {
		tempLen := idx + 1
		if minLen < tempLen {
			return
		}

		for _, w := range transMap[word] {
			if isUsed[w] {
				continue
			}

			path[idx] = w
			if w == endWord {
				temp := make([]string, tempLen)
				copy(temp, path)

				if minLen > tempLen {
					minLen = tempLen
					res = [][]string{temp}
				} else {
					// minLen == tempLen
					res = append(res, temp)
				}
				return
			}
			isUsed[w] = true
			dfs(w, path, idx+1)
			isUsed[w] = false
		}
	}

	path := make([]string, len(words)*2)
	path[0] = beginWord
	dfs(beginWord, path, 1)

	return res
}

func isTransable(a, b string) bool {
	diff := false
	for i := range a {
		if a[i] != b[i] {
			if !diff {
				diff = true
			} else {
				return false
			}
		}
	}

	return true
}
