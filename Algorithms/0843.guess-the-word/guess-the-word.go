package problem0843

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// Master is the Master's API interface.
// You should not implement it, or speculate about its implementation
// type Master struct {}
// func (this *Master) Guess(word string) int {}
type Master = kit.Master

// 解答参考 https://leetcode.com/problems/guess-the-word/discuss/133862/Random-Guess-and-Minimax-Guess-with-Comparison

func findSecretWord(wordList []string, master *Master) {
	matches := 0
	// 题目默认是猜 10 次
	for i := 0; i < 10; i++ {
		// count[w] 代表了 w 与 wordList 中单词的匹配程度
		// 数值越高，越匹配
		count := make(map[string]int, len(wordList))
		for _, w := range wordList {
			for _, b := range wordList {
				count[w] += match(w, b)
			}
		}

		key := ""
		max := 0
		for _, w := range wordList {
			if max < count[w] {
				max = count[w]
				key = w
			}
		}
		// 现在 key 与 wordList 中别的单词最具有相似性

		matches = master.Guess(key)
		// wordList 中的单词长度都为 6
		if matches == 6 {
			// 猜到了
			return
		}

		// 没有猜到就缩小 wordList 的范围
		newList := make([]string, 0, len(wordList))
		for _, w := range wordList {
			if match(key, w) == matches {
				// 因为 matches = match(key, secret)
				// 所以，符合 match(key, w) == matches 的 w 才有可能是 secret
				newList = append(newList, w)
			}
		}

		wordList = newList
	}
}

// a, b 总是一样长
func match(a, b string) int {
	res := 0
	size := len(a)
	for i := 0; i < size; i++ {
		if a[i] == b[i] {
			res++
		}
	}
	return res
}
