package Problem0722

import "strings"

func removeComments(source []string) []string {
	res := dealSingle(source)
	res = dealBlock(res)
	return res
}

func dealSingle(source []string) []string {
	res := make([]string, 0, len(source))
	for _, s := range source {
		i := strings.Index(s, "//")
		if i == -1 {
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
	for _, s := range source {
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
		}
	}

	return res
}
