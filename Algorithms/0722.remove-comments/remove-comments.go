package Problem0722

import "strings"

// NOTICE: 使用正则表达式替换，可以很简单地解决这个题目

func removeComments(source []string) []string {
	s := strings.Join(source, "$$") + "$$"
	s = helper(s)
	source = split(s)
	return source
}

func helper(s string) string {
	i := strings.Index(s, "//")
	j := strings.Index(s, "/*")

	if (i == -1 && 0 <= j) ||
		(0 <= j && j < i) {
		k := j + strings.Index(s[j:], "*/")
		return s[:j] + helper(s[k+2:])
	}

	if (j == -1 && 0 <= i) ||
		(0 <= i && i < j) {
		k := i + strings.Index(s[i:], "$$")
		return s[:i] + helper(s[k:])
	}

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

func dealLine(source []string) []string {
	res := make([]string, 0, len(source))
	for _, s := range source {
		i := strings.Index(s, "//")
		j := strings.Index(s, "/*")
		k := strings.LastIndex(s, "*/")

		if i == 0 && k < 0 { // 后面要求没有 "*/"
			continue
		}

		if i == -1 || // 无 "//"
			(0 <= j && j < i) || // "/*" 在 "//" 前面起作用了
			i < k { // "//" 后面还有 "*/" 所以要留着
			res = append(res, s)
			continue
		}

		res = append(res, s[:i])
	}
	return res
}

func dealBlock(source []string) []string {
	res := make([]string, 0, len(source))
	isInBlock := false
	temp := ""
	for k := 0; k < len(source); k++ {
		s := source[k]
		if isInBlock {
			i := strings.LastIndex(s, "*/")
			if i >= 0 {
				temp += s[i+2:]
				if temp != "" {
					res = append(res, temp)
				}
				temp = ""
				isInBlock = false
			}
			continue
		}

		i := strings.Index(s, "/*")

		if i == -1 {
			res = append(res, s)
		} else {
			isInBlock = true
			temp = s[:i]
			k--
		}
	}

	return res
}
