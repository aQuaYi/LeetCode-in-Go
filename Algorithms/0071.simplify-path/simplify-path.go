package problem0071

import "strings"

func simplifyPath(path string) string {
	lp := len(path)
	stack := make([]string, 0, lp/2)
	dir := make([]byte, 0, lp)

	for i := 0; i < lp; i++ {
		// 使用前，清空 dir
		// 这个方法比 dir = []byte{} 快了近 8 倍
		dir = dir[:0]
		// 获取 dir
		for i < lp && path[i] != '/' {
			dir = append(dir, path[i])
			i++
		}

		s := string(dir)

		switch s {
		case ".", "":
			// do nothing
		case "..":
			if len(stack) > 0 {
				// pop
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, s)
		}
	}

	return "/" + strings.Join(stack, "/")
}
