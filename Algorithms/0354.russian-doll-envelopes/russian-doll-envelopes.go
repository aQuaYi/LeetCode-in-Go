package problem0354

import "sort"

func maxEnvelopes(e [][]int) int {
	if len(e) <= 1 {
		return len(e)
	}

	sort.Sort(sortedEnvelopes(e))

	var tails []int

	for i := 0; i < len(e); i++ {
		lo, hi := 0, len(tails)
		// 按照信封高度来二分查找插入
		for lo < hi {
			mid := (lo + hi) / 2
			if e[i][1] <= tails[mid] {
				hi = mid
			} else {
				lo = mid + 1
			}
		}

		// 出现了新高度，添加到 tail 尾部
		if lo == len(tails) {
			tails = append(tails, e[i][1])
		} else {
			// e[i][1] 是一个更合理的高度，
			// 用 e[i][1] 替换掉 tails[lo] 可以避免遗漏最优解
			tails[lo] = e[i][1]
		}
	}

	return len(tails)
}

type sortedEnvelopes [][]int

func (s sortedEnvelopes) Len() int {
	return len(s)
}

func (s sortedEnvelopes) Less(i, j int) bool {
	if s[i][0] == s[j][0] {
		// 同样宽度时，矮信封排在后面
		return s[i][1] > s[j][1]
	}
	return s[i][0] < s[j][0]
}

func (s sortedEnvelopes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
