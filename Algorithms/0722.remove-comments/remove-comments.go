package problem0722

import "strings"

// TODO: 使用正则表达式替换，可以很简单地解决这个题目

func removeComments(source []string) []string {
	// 末尾添加 "$$" 是为了当最后一行是 "// ..." 时，能找到终止符号
	s := strings.Join(source, "$$") + "$$"
	s = helper(s)
	source = split(s)
	return source
}

func helper(s string) string {
	i := strings.Index(s, "/*")
	j := strings.Index(s, "//")

	// "/*" 起作用了
	if 0 <= i &&
		(j == -1 || i < j) {
		// i+2 是为了避免匹配到 "/*/" 中的 "*/"
		k := i + 2 + strings.Index(s[i+2:], "*/")
		return s[:i] + helper(s[k+2:]) // 删除了 s[i:k+2]
	}

	// "//" 起作用了
	if 0 <= j &&
		(i == -1 || j < i) {
		k := j + strings.Index(s[j:], "$$")
		return s[:j] + helper(s[k:]) // 删除了 s[j:k]
	}

	// 已经没有注释了
	return s
}

func split(s string) []string {
	ss := strings.Split(s, "$$")
	res := make([]string, 0, len(ss))
	for _, s := range ss {
		if s != "" {
			res = append(res, s)
		}
	}
	return res
}
