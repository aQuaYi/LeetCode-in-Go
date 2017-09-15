package Problem0126

func findLadders(beginWord string, endWord string, words []string) [][]string {
	// 因为 beginWord 不能做 transformed word 很碍事，就先删掉，好让后面的逻辑简明一些
	words = deleteBeginWord(words, beginWord)

	// trans 用来记录 k -> v[i] 的转换关系。
	trans := map[string][]string{}
	// isMatchedEndWord 用于在生成 trans 的过程中，存在 -> endWord 的转换关系
	// 用于提前结束
	isMatchedEndWord := false
	// cnt 用于记录生成trans的迭代次数
	// 其实也是最短路径的长度
	cnt := 1
	var bfs func([]string, []string)
	// 使用 bfs 方法，递归地生成 trans
	bfs = func(words, nodes []string) {
		cnt++
		// words 中的 w
		//     与 nodes 中的 n -> w，在 w 会被放入 newNodes
		//     否则，w 会被放入 newWords
		newWords := make([]string, 0, len(words))
		newNodes := make([]string, 0, len(words))
		for _, w := range words {
			isMatched := false
			for _, n := range nodes {
				if isTransable(w, n) {
					isMatched = true
					trans[n] = append(trans[n], w)
				}
			}
			if isMatched {
				newNodes = append(newNodes, w)
				if w == endWord {
					isMatchedEndWord = true
				}
			} else {
				newWords = append(newWords, w)
			}
		}

		if isMatchedEndWord || // 匹配到 endWord 说明已经找到了所有的最短路径
			len(newWords) == 0 || // 全部匹配完成
			len(newNodes) == 0 { // newWords 中单词，是 beginWord 无法 trans 到的
			return
		}

		// 继续完善 trans
		bfs(newWords, newNodes)
	}

	nodes := []string{beginWord}
	bfs(words, nodes)

	res := [][]string{}
	if !isMatchedEndWord {
		// beginWord 无法 trans 到 endWord
		return res
	}

	var dfs func([]string, int)
	dfs = func(path []string, idx int) {
		if idx == cnt {
			if path[idx-1] == endWord {
				res = append(res, deepCopy(path))
			}
			return
		}

		word := path[idx-1]
		for _, w := range trans[word] {
			path[idx] = w
			dfs(path, idx+1)
		}
	}
	path := make([]string, cnt)
	path[0] = beginWord
	dfs(path, 1)

	return res
}

func deepCopy(src []string) []string {
	temp := make([]string, len(src))
	copy(temp, src)
	return temp
}

func deleteBeginWord(words []string, beginWord string) []string {
	i := 0
	for ; i < len(words); i++ {
		if words[i] == beginWord {
			break
		}
	}

	if i == len(words) {
		return words
	}

	return append(words[:i], words[i+1:]...)
}

func isTransable(a, b string) bool {
	// onceAgain == true 说明已经出现过不同的字符了
	onceAgain := false
	for i := range a {
		if a[i] != b[i] {
			if onceAgain {
				return false
			}
			onceAgain = true
		}
	}

	return true
}
