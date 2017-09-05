package Problem0071

import "strings"

func simplifyPath(path string) string {
	lp := len(path)
	stack := []string{}
	dir := []byte{}

	for i := 0; i < lp; i++ {
		// 截取 dir
		for i < lp && path[i] != '/' {
			dir = append(dir, path[i])
			i++
		}

		s := string(dir)
		switch s {
		case ".", "":
			continue
		case "..":
			if len(stack) > 0 {
				// pop
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, s)

		}

		dir = dir[:0]
	}

	return "/" + strings.Join(stack, "/")
}
