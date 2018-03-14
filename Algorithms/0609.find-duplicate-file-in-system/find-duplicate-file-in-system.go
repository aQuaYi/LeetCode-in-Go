package problem0609

import (
	"strings"
)

func findDuplicate(path []string) [][]string {
	// m 用于记录拥有相同文件内容的各个文件的个数
	m := make(map[string][]string, len(path)*2)
	for i := 0; i < len(path); i++ {
		analyze(path[i], m)
	}

	res := make([][]string, 0, len(m))
	for _, files := range m {
		if len(files) > 1 {
			res = append(res, files)
		}
	}

	return res
}

// 把 path 中的文件，按照文件内容，归类记录
func analyze(path string, m map[string][]string) {
	fs := strings.Split(path, " ")
	p := fs[0]
	for i := 1; i < len(fs); i++ {
		f, c := fileAndContent(fs[i])
		m[c] = append(m[c], p+"/"+f)
	}
}

// 形为 filename(fileContent) 的 s 分解为 filename 和 fileContent 返回
func fileAndContent(s string) (string, string) {
	i := strings.IndexByte(s, '(')
	return s[:i], s[i+1 : len(s)-1]
}
