package problem0884

import (
	"sort"
	"strings"
)

func uncommonFromSentences(A string, B string) []string {
	tmp := strings.Split(A, " ")
	tmp = append(tmp, strings.Split(B, " ")...)

	// 排序完成后， "" 会在首， "~" 会在尾
	// 安排这两个不相干的单词进入 tmp ，有利于简化后面的判断逻辑
	tmp = append(tmp, "", "~")
	sort.Strings(tmp)

	size := len(tmp)

	res := make([]string, 0, size)

	for i := 1; i+1 < size; i++ {
		if tmp[i-1] != tmp[i] && tmp[i] != tmp[i+1] {
			res = append(res, tmp[i])
		}
	}

	return res
}
