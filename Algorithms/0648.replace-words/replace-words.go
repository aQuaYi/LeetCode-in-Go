package problem0648

import (
	"sort"
	"strings"
)

func replaceWords(dict []string, sentence string) string {
	// hasRoot 收集了所有的 root
	hasRoot := make(map[string]bool, len(dict))
	hasLen := make([]bool, 101)
	for i := range dict {
		hasRoot[dict[i]] = true
		hasLen[len(dict[i])] = true
	}

	// ls 收集了所有 root 的长度
	ls := make([]int, 0, 101)
	for i, ok := range hasLen {
		if ok {
			ls = append(ls, i)
		}
	}
	sort.Ints(ls)

	words := strings.Split(sentence, " ")
	for i, w := range words {
		for _, l := range ls {
			// w 的前 l 个字符是字根就修改 w 并结束 for 循环
			if l < len(w) && hasRoot[w[:l]] {
				words[i] = w[:l]
				break
			}
		}
	}

	return strings.Join(words, " ")
}
