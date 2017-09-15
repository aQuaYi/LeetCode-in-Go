package Problem0127

func ladderLength(beginWord string, endWord string, words []string) int {

	// isTransedEndWord 用于在生成 trans 的过程中标记，存在 ->endWord 的转换关系
	// 用于提前结束
	isTransedEndWord := false
	// cnt 用于记录生成trans的迭代次数
	// 其实也是最短路径的长度
	res := 1
	var bfs func([]string, []string)
	// 使用 bfs 方法，递归地生成 trans
	bfs = func(words, nodes []string) {
		res++
		// words 中的 w
		//     与 nodes 中的 n ，可以实现 n->w 的转换，
		//       则，w 会被放入 newNodes
		//     否则，w 会被放入 newWords
		newWords := make([]string, 0, len(words))
		newNodes := make([]string, 0, len(words))
		for _, w := range words {
			isTransed := false
			for _, n := range nodes {
				if isTransable(n, w) {
					isTransed = true

				}
			}

			if isTransed {
				if w == endWord {
					isTransedEndWord = true
					return
				}
				newNodes = append(newNodes, w)
			} else {
				newWords = append(newWords, w)
			}
		}

		if len(newWords) == 0 || // words 的所有单词都已经可以从 beginWord trans 到
			len(newNodes) == 0 { // newWords 中单词，是 beginWord 无法 trans 到的
			return
		}

		// 上面没有 return，要继续完善 trans
		bfs(newWords, newNodes)
	}

	// 第一代 nodes 含有且仅含有 beginWord
	nodes := []string{beginWord}
	bfs(words, nodes)

	if !isTransedEndWord {
		// beginWord 无法 trans 到 endWord
		return 0
	}

	return res
}

func isTransable(a, b string) bool {
	if a == b {
		return false
	}
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
