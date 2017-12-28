package Problem0722

import "strings"

// NOTICE: 使用正则表达式替换，可以很简单地解决这个题目

func removeComments(source []string) []string {
	source = dealLine(source)
	source = dealBlock(source)
	return source
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
			j < i || // "/*" 在 "//" 前面起作用了
			i < k { // "//" 后面还有 "*/" 所以要留着
			res = append(res, s)
			continue
		}

		res = append(res, s[i:])
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
