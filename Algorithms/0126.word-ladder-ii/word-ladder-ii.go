package Problem0126

import (
	"fmt"
)

func findLadders(beginWord string, endWord string, words []string) [][]string {
	res := [][]string{}
	// 因为 beginWord 不能做 transformed word
	// 很碍事，就先删掉，好让后面的逻辑简明一些
	words = deleteBeginWord(words, beginWord)

	trans := map[string][]string{}
	isMatchedEndWord := false
	var bfs func([]string, []string)

	bfs = func(words, nodes []string) {
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

		if isMatchedEndWord || len(newWords) == 0 || len(newNodes) == 0 {
			return
		}

		bfs(newWords, newNodes)
	}

	nodes := make([]string, len(words)*2)
	nodes[0] = beginWord
	bfs(words, nodes)

	fmt.Println(trans)
	return res
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
