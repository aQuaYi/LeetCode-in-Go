package Problem0076

import (
	"sort"
)

func minWindow(s string, t string) string {
	rec := make(map[byte][]int, len(s))

	for i := range s {
		rec[s[i]] = append(rec[s[i]], i)
	}

	// 检查，t 中的字母是否都在 s 中
	for i := range t {
		if len(rec[t[i]]) == 0 {
			return ""
		}
	}

	lt := len(t)
	min := len(s) + 1
	first, last := 0, 0
	list := make([]int, lt)

	var dfs func(int)
	dfs = func(idx int) {
		if idx == lt {
			temp := deepCopy(list)
			sort.Ints(temp)
			dist := temp[lt-1] - temp[0]
			if min > dist {
				first, last = temp[0], temp[lt-1]
				min = dist
			}
			return
		}

		indexs := deepCopy(rec[t[idx]])

		for i := 0; i < len(indexs); i++ {
			list[idx] = indexs[i]

			dfs(idx + 1)
		}
	}

	dfs(0)

	return s[first : last+1]
}

func deepCopy(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}
