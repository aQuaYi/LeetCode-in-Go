package Problem0522

func findLUSlength(strs []string) int {
	count := make(map[string]int, len(strs))
	for _, s := range strs {
		count[s]++
	}

	maxLen := -1
	for s, c := range count {
		if c == 1 && maxLen < len(s) {
			maxLen = len(s)
		}
	}

	return maxLen
}

// 如果 s 是 ss 中某一个比 s 长的字符串的子字符串，返回 true
// ss 是排序过的
func isSubOf(s string, ss []string) bool {

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
