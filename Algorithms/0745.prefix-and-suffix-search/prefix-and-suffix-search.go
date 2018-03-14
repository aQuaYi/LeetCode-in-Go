package problem0745

// WordFilter 是字符过滤器
type WordFilter struct {
	ords map[string]int
}

// Constructor 返回 WordFilter
func Constructor(words []string) WordFilter {
	o := make(map[string]int, len(words)*5)
	for k := 0; k < len(words); k++ {
		for i := 0; i <= 10 && i <= len(words[k]); i++ {
			for j := len(words[k]); 0 <= j && len(words[k])-10 <= j; j-- {
				pps := words[k][:i] + "#" + words[k][j:]
				o[pps] = k
			}
		}
	}
	return WordFilter{ords: o}
}

// F 用于查找同时拥有前缀和后缀的单词的索引号，同时存在多个符合条件的单词的话，返回其中的最大值
func (w *WordFilter) F(prefix, suffix string) int {
	pps := prefix + "#" + suffix
	if index, ok := w.ords[pps]; ok {
		return index
	}
	return -1
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
