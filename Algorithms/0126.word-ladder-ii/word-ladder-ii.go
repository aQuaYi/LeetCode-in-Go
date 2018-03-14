package problem0126

func findLadders(beginWord string, endWord string, words []string) [][]string {
	// 因为 beginWord 不能做 transformed word
	// 先删掉 words 中的 beginWord，
	// 可以让后面的 trans 少很多的断头 path，会加快程序。
	// 也更符合题意
	// 删除下面这句，程序也能 accepted，但是会从 269ms 减慢到 319ms
	words = deleteBeginWord(words, beginWord)

	// trans 用来查找 k->？
	trans := map[string][]string{}
	// isTransedEndWord 用于在生成 trans 的过程中标记，存在 ->endWord 的转换关系
	// 用于提前结束
	isTransedEndWord := false
	// cnt 用于记录生成trans的迭代次数
	// 其实也是最短路径的长度
	cnt := 1
	var bfs func([]string, []string)
	// 使用 bfs 方法，递归地生成 trans
	bfs = func(words, nodes []string) {
		cnt++
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
					trans[n] = append(trans[n], w)
					isTransed = true
				}
			}

			if isTransed {
				newNodes = append(newNodes, w)
				if w == endWord {
					isTransedEndWord = true
				}
			} else {
				newWords = append(newWords, w)
			}
		}

		if isTransedEndWord || // 转换到了 endWord 说明已经找到了所有的最短路径
			len(newWords) == 0 || // words 的所有单词都已经可以从 beginWord trans 到
			len(newNodes) == 0 { // newWords 中单词，是 beginWord 无法 trans 到的
			return
		}

		// 上面没有 return，要继续完善 trans
		bfs(newWords, newNodes)
	}

	// 第一代 nodes 含有且仅含有 beginWord
	nodes := []string{beginWord}
	bfs(words, nodes)

	res := [][]string{}
	if !isTransedEndWord {
		// beginWord 无法 trans 到 endWord
		return res
	}

	path := make([]string, cnt)
	path[0] = beginWord

	var dfs func(int)
	// 使用 dfs 方法，生成最短路径
	dfs = func(idx int) {
		if idx == cnt {
			// path 已经填充完毕
			if path[idx-1] == endWord {
				// 最后一个单词是 endWord，说明这是一条最短路径
				res = append(res, deepCopy(path))
			}
			return
		}

		prev := path[idx-1]
		for _, w := range trans[prev] {
			// 利用 prev->w 填充 path[idx]
			path[idx] = w
			dfs(idx + 1)
		}
	}

	dfs(1)

	return res
}

func deepCopy(src []string) []string {
	temp := make([]string, len(src))
	copy(temp, src)
	return temp
}

// 题目中说了，words 中没有重复的单词，
// 所以，beginWord 最多出现一次
func deleteBeginWord(words []string, beginWord string) []string {
	i, size := 0, len(words)
	for ; i < size; i++ {
		if words[i] == beginWord {
			break
		}
	}

	if i == size {
		return words
	}

	words[i] = words[size-1]
	return words[:size-1]
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
