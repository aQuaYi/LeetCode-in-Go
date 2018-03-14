package problem0524

import (
	"sort"
)

func findLongestWord(s string, d []string) string {
	// 排序后，第一个符合条件的字符串，就是答案
	sort.Sort(stringSlice(d))
	for i := range d {
		if isSub(s, d[i]) {
			return d[i]
		}
	}
	return ""
}

// stringSlice 实现了 sort.Interface 接口
type stringSlice []string

func (s stringSlice) Len() int { return len(s) }

func (s stringSlice) Less(i, j int) bool {
	if len(s[i]) == len(s[j]) {
		return s[i] < s[j]
	}
	return len(s[i]) > len(s[j])
}

func (s stringSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// return true if sub is sub of s
func isSub(s, sub string) bool {
	if len(s) < len(sub) {
		return false
	}
	i, j := 0, 0
	for i < len(s) && j < len(sub) {
		if s[i] == sub[j] {
			j++
		}
		i++
	}
	return j == len(sub)
}
