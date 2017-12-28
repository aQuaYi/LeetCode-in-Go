package Problem0722

import "strings"

// NOTICE: 使用正则表达式替换，可以很简单地解决这个题目

func removeComments(source []string) []string {
	res := dealSingle(source)
	res = dealBlock(res)
	return res
}

func dealSingle(source []string) []string {
	res := make([]string, 0, len(source))
	for _, s := range source {
		i := strings.Index(s, "//")
		j := strings.Index(s, "/*")
		k := strings.Index(s, "*/")

		if i == -1 ||
			j >= 0 ||
			k >= 0 {
			res = append(res, s)
		} else if i > 0 {
			res = append(res, s[:i])
		}
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
