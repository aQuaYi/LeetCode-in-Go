package Problem0331

import (
	"strings"
)

func isValidSerialization(preorder string) bool {
	if preorder == "#" {
		return true
	}

	ss := strings.Split(preorder, ",")

	if (len(ss)-3)%2 != 0 {
		return false
	}

	return check(ss)
}

func check(ss []string) bool {
	n := len(ss)

	if n < 3 {
		return false
	}

	if n == 3 {
		if ss[0] != "#" &&
			ss[1] == "#" &&
			ss[2] == "#" {
			return true
		}
		return false
	}

	temp := make([]string, 0, n)

	var i = 0
	for i < n {
		if i+2 < n && ss[i] != "#" && ss[i+1] == "#" && ss[i+2] == "#" {
			temp = append(temp, "#")
			i += 3
			continue
		}

		temp = append(temp, ss[i])
		i++
	}

	return check(temp)
}
