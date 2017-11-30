package Problem0522

import (
	"sort"
)

func findLUSlength(strs []string) int {
	sort.Sort(stringSlice(strs))

	for i := 0; i < len(strs); i++ {
		for i+1 < len(strs) && strs[i] == strs[i+1] {
			i++
		}

		if i < len(strs) && i-1 >= 0 && !isSub(strs[i], strs[i-1]) {
			return len(strs[i])
		}
	}

	return -1
}

// 如果 a 是 b 的子字符串，返回 true
func isSub(a, b string) bool {
	aLen, bLen := len(a), len(b)
	i, j := 0, 0

	for i < aLen && j < bLen {
		if a[i] == b[j] {
			i++
		}
		j++
	}

	return i == aLen
}

// strings 实现了 sort.Interface 接口
type stringSlice []string

func (ss stringSlice) Len() int { return len(ss) }

func (ss stringSlice) Less(i, j int) bool {
	if len(ss[i]) == len(ss[j]) {
		return ss[i] > ss[j]
	}
	return len(ss[i]) > len(ss[j])
}

func (ss stringSlice) Swap(i, j int) { ss[i], ss[j] = ss[j], ss[i] }
