package problem0522

import (
	"sort"
)

func findLUSlength(strs []string) int {
	// 统计每个单词出现的次数
	count := make(map[string]int, len(strs))
	for _, s := range strs {
		count[s]++
	}

	// 让 strs 中的每个单词只出现一次
	// 这样的话，后面检查是否为子字符串的时候，可以避免重复检查
	strs = strs[:0]
	for s := range count {
		strs = append(strs, s)
	}

	// 按照字符串的长度进行降序排列
	// 后面检查是否为子字符串的时候，不用每个都检查一遍
	sort.Sort(stringSlice(strs))

	for i, s := range strs {
		// 筛掉重复的字符串
		if count[s] > 1 {
			continue
		}
		// s 必须不是比 Ta 长的字符串的子字符串才能是要找的解答
		if !isSubOf(s, strs[:i]) {
			return len(s)
		}
	}

	return -1
}

// 如果 s 是 ss 中某一个比 s 长的字符串的子字符串，返回 true
// ss 是排序过的
func isSubOf(s string, ss []string) bool {
	for i := range ss {
		if isSub(s, ss[i]) {
			return true
		}
	}
	return false
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

// stringSlice 实现了 sort.Interface 接口
type stringSlice []string

func (ss stringSlice) Len() int { return len(ss) }

func (ss stringSlice) Less(i, j int) bool { return len(ss[i]) > len(ss[j]) }

func (ss stringSlice) Swap(i, j int) { ss[i], ss[j] = ss[j], ss[i] }
