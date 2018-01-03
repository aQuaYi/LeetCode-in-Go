package Problem0746

import (
	"strings"
)

// WordFilter 是字符过滤器
type WordFilter struct {
	ords []string
}

// Constructor 返回 WordFilter
func Constructor(words []string) WordFilter {
	return WordFilter{ords: words}
}

// F 用于查找同时拥有前缀和后缀的单词的索引号，同时存在多个符合条件的单词的话，返回其中的最大值
func (w *WordFilter) F(prefix, suffix string) int {
	n := len(w.ords)

	for i := n - 1; 0 <= i; i-- {
		if strings.HasPrefix(w.ords[i], prefix) &&
			strings.HasSuffix(w.ords[i], suffix) {
			return i
		}
	}

	return -1
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
