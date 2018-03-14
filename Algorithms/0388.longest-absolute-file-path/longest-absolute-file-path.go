package problem0388

import (
	"strings"
)

func lengthLongestPath(path string) int {
	if !hasFile(path) {
		return 0
	}

	// 添加了 "\n" 这个字符，所以，最后的答案需要 -1
	return llp("\n"+path) - 1
}

func llp(path string) int {
	if !hasFile(path) {
		// 按照题意，最长的路径中，必须包含文件
		return 0
	}

	if !hasSub(path) {
		// 没有子文件夹，又是一个文件
		return len(path)
	}

	var i = nextCR(path, -1)
	// path[:i] 是 父目录
	dirLen := i

	var j, maxSub int
	for i < len(path) {
		j = nextCR(path, i)
		// path[i:j] 是 子目录
		// 替换缩进层级后，可以递归处理
		maxSub = max(maxSub, llp(strings.Replace(path[i+1:j], "\n\t", "\n", -1)))
		i = j
	}

	// +1 是路径分隔符 "/" 的长度
	return dirLen + 1 + maxSub
}

func hasFile(path string) bool {
	return strings.Contains(path, ".")
}

// 有下级，返回 true
func hasSub(path string) bool {
	return strings.Contains(path, "\n")
}

func nextCR(path string, idx int) int {
	var i int
	for i = idx + 1; i < len(path); i++ {
		if path[i:i+1] == "\n" &&
			(i+1 == len(path) || // 防止路径末尾存在换行的情况，比如 "a.txt\n"
				path[i+1:i+2] != "\t") {
			break
		}
	}
	return i
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
