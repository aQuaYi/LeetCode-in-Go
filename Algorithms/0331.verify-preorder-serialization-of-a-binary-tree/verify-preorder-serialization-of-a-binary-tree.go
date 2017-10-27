package Problem0331

import (
	"strings"
)

func isValidSerialization(preorder string) bool {
	ss := strings.Split(preorder, ",")

	if (len(ss)-3)%2 != 0 {
		return false
	}

	return check(ss)
}

func check(ss []string) bool {
	n := len(ss)

	if n == 1 && ss[0] == "#" {
		return true
	}

	temp := make([]string, 0, n)

	var i = 0
	for i < n {
		if i+2 < n && ss[i] != "#" && ss[i+1] == "#" && ss[i+2] == "#" {
			// ss[i] 的两个子节点都是 叶子 节点
			// 把 ss[i] 变成 叶子 节点，删除其子节点
			temp = append(temp, "#")
			i += 3
			continue
		}

		temp = append(temp, ss[i])
		i++
	}

	return len(temp) < n && check(temp)
}
