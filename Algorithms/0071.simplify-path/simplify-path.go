package Problem0071

import (
	"strings"
)

func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	res := []string{}

	for i, p := range paths {
		switch p {
		case ".", "":
			continue
		case "..":
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		default:
			res = append(res, paths[i])
		}
	}

	return "/" + strings.Join(res, "/")
}
